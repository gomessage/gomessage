#!/usr/bin/env bash

VERSION="1.0.10"
OUTPUT_PATH="./build/${VERSION}/"
GOOS_LIST=(darwin linux windows)
GOARCH_LIST=(amd64)
EXRS='^build|^bak|^docs|^swagger|^deploy|^LICENSE|.sh$|.sqlite3$|.bak$|.md$|prof$'

build() {
    #yum groupinstall -y "Development Tools"
    go get github.com/beego/bee/v2@v2.0.2

    #去掉不必要的包
    go mod tidy

    #编译之前：删除不必要的东西，不确定是否会被编译进去，但是以防万一比较好
    rm -rf ./*.log
    rm -rf ./*.tar.gz
    rm -rf ./GoMessage
    rm -rf ./gomessage-*

    #修改runmode参数，让编译后的应用程序运行在prod环境中
    gsed -i '/^runmode/c runmode = prod' ./conf/app.conf

    mkdir -p ${OUTPUT_PATH}${baoName}

    if [[ ${GOOS} == "linux" ]]; then
        bee pack -a gomessage \
            -o "${OUTPUT_PATH}${baoName}/" \
            -exr ${EXRS} \
            -be CGO_ENABLED=1 \
            -be CGO_LDFLAGS="-static" \
            -be CC=x86_64-linux-musl-gcc \
            -be CXX=x86_64-linux-musl-g++ \
            -be GOOS="${GOOS}" \
            -be GOARCH="${GOARCH}"

        cp ./deploy/install.sh ${OUTPUT_PATH}${baoName}/
        cp ./deploy/uninstall.sh ${OUTPUT_PATH}${baoName}/
        #cp ./deploy/Dockerfile ${OUTPUT_PATH}${baoName}/

    elif [[ ${GOOS} == "windows" ]]; then
        bee pack -a gomessage \
            -o "${OUTPUT_PATH}${baoName}/" \
            -exr ${EXRS} \
            -be CGO_ENABLED=1 \
            -be CGO_CFLAGS="-g -O2 -Wno-return-local-addr" \
            -be CC=x86_64-w64-mingw32-gcc \
            -be CXX=x86_64-w64-mingw32-g++ \
            -be GOOS="${GOOS}" \
            -be GOARCH="${GOARCH}"

    elif [[ ${GOOS} == "darwin" ]]; then
        bee pack -a gomessage \
            -o "${OUTPUT_PATH}${baoName}/" \
            -exr ${EXRS} \
            -be CGO_ENABLED=1 \
            -be GOOS="${GOOS}" \
            -be GOARCH="${GOARCH}"
    else
        echo "GOOS版本未识别~"
    fi

    tar -zcvf ${OUTPUT_PATH}${baoName}.tar.gz -C ${OUTPUT_PATH} ${baoName}

    #编译完成之后再把runmode参数修改回来，这样下次都不用手动修改了
    gsed -i '/^runmode/c runmode = dev' ./conf/app.conf

}

#清空指定目录下的内容（只执行一次）
mkdir -p ${OUTPUT_PATH}
echo ${OUTPUT_PATH}

#rm -rf "${OUTPUT_PATH}"*
rm -rf ./build/*
echo "历史包清理完成..."

#循环编译指定版本的包
for GOOS in "${GOOS_LIST[@]}"; do
    for GOARCH in "${GOARCH_LIST[@]}"; do
        #遍历数组后拿到的<临时变量>拼接成<新的包名>
        baoName="gomessage-${VERSION}-${GOOS}-${GOARCH}"
        build
    done
done
