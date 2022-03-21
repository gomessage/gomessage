#!/usr/bin/env bash


VERSION="1.0.9"  #编译之后的包版本（只修改这一个就行了，下文别的东西不要动）


#下面这些参数尽量不要动（重复通读整个脚本，明白这些变量的用途，否则不要进行修改和变动）
OUTPUT_PATH="./build/${VERSION}/"  #制品输出路径
GOOS_LIST=(darwin linux windows)  #系统平台
GOARCH_LIST=(amd64)  #系统脚骨
EXRS='^build|^bak|^docs|^swagger|^deploy|^LICENSE|.sh$|.sqlite3$|.bak$|.md$|prof$'  #编译时要忽略的文件


#该函数负责实际的代码交叉编译
build() {
    #bee工具的安装
    go get -u github.com/beego/bee/v2

    #去掉不必要的包
    go mod tidy

    #编译之前：删除不必要的东西，不确定是否会被编译进去，但是以防万一比较好
    rm -rf ./*.log
    rm -rf ./*.tar.gz
    rm -rf ./GoMessage

    #修改runmode参数，让编译后的应用程序运行在prod环境中
    gsed -i '/^runmode/c runmode = prod' ./conf/app.conf

    mkdir -p ${OUTPUT_PATH}"${baoName}"

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
        # cp ./deploy/Dockerfile ${OUTPUT_PATH}${baoName}/   #不再复制dockerfile文件进入最终的安装包中，次命令留作备忘

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


#一定要确定build目录下是干净的
#rm -rf "${OUTPUT_PATH}"*
rm -rf ./build/*
echo "历史包清理完成..."


#开始进行编译：循环遍历两个参数：系统平台、系统架构；然后进行分别编译
for GOOS in "${GOOS_LIST[@]}"; do
    for GOARCH in "${GOARCH_LIST[@]}"; do
        #遍历数组后拿到的<临时变量>拼接成<新的包名>
        baoName="gomessage-${VERSION}-${GOOS}-${GOARCH}"
        build
    done
done
