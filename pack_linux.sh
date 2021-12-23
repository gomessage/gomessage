#!/usr/bin/env bash

VERSION="v1.0.5"
OUTPUT_PATH="./build/${VERSION}/"
#GOOS_LIST=(darwin linux windows)
GOOS_LIST=(linux)
GOARCH_LIST=(amd64)

build() {
    yum groupinstall -y "Development Tools"
    go get -u github.com/beego/bee/v2

    #去掉不必要的包
    go mod tidy

    #编译之前：删除不必要的东西，不确定是否会被编译进去，但是以防万一比较好
    rm -rf ./*.log
    rm -rf ./*.tar.gz
    rm -rf ./GoMessage

    #修改runmode参数，让编译后的应用程序运行在prod环境中
    sed -i '/^runmode/c runmode = prod' ./conf/app.conf

    mkdir -p ${OUTPUT_PATH}${GOOS}

    if [[ ${GOOS} == "linux" ]]; then
        bee pack -a gomessage \
            -o "${OUTPUT_PATH}${GOOS}/" \
            -exr "^pack.sh|^build|README.md|.sh\$|bak|.sqlite3\$|LICENSE|bak\$|docs" \
            -be CGO_ENABLED=1 \
            -be GOOS="${GOOS}" \
            -be GOARCH="${GOARCH}"

        cp ./install.sh ${OUTPUT_PATH}${GOOS}/
        cp ./uninstall.sh ${OUTPUT_PATH}${GOOS}/
    else
        bee pack -a gomessage \
            -o "${OUTPUT_PATH}${GOOS}/" \
            -exr "^pack.sh|^build|README.md|.sh\$|bak|.sqlite3\$|LICENSE|bak\$" \
            -be CGO_ENABLED=1 \
            -be GOOS="${GOOS}" \
            -be GOARCH="${GOARCH}"
    fi

    tar -zcvf ${OUTPUT_PATH}${baoName}.tar.gz -C ${OUTPUT_PATH}${GOOS}/ .

    #编译完成之后再把runmode参数修改回来，这样下次都不用手动修改了
    sed -i '/^runmode/c runmode = dev' ./conf/app.conf

}

#清空指定目录下的内容（只执行一次）
mkdir -p ${OUTPUT_PATH}
echo ${OUTPUT_PATH}
rm -rf "${OUTPUT_PATH}"*
echo "历史包清理完成..."

#循环编译指定版本的包
for GOOS in "${GOOS_LIST[@]}"; do
    for GOARCH in "${GOARCH_LIST[@]}"; do
        baoName="gomessage-${VERSION}-${GOOS}-${GOARCH}"
        build
    done
done
