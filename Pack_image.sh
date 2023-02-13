#!/usr/bin/env bash

VERSION="1.0.11"
OUTPUT_PATH="./build/${VERSION}/"
GOOS_LIST=(darwin linux windows)
GOARCH_LIST=(amd64)

build() {

    #创建工作目录
    mkdir -p ${OUTPUT_PATH}${baoName}

    if [[ ${GOOS} == "linux" ]]; then

        #docker镜像编译（如果本地docker服务开启，则自动编译image，如果本地docker服务未运行，则跳过该步骤）
        if (launchctl list | grep com.docker.docker); then
            if [[ $(launchctl list | grep com.docker.docker | awk '{print $1}') == "-" ]]; then
                echo "Docker服务进程号为-，因此跳过编译docker镜像~"
            else
                #复制dockerfile到安装包目录
                #cp ./deploy/Dockerfile ${OUTPUT_PATH}${baoName}/

                #编译对应版本的image
                docker build -t taycc/gomessage:latest -f ./deploy/Dockerfile "${OUTPUT_PATH}${baoName}/"
                docker build -t taycc/gomessage:${VERSION} -f ./deploy/Dockerfile "${OUTPUT_PATH}${baoName}/"
            fi
        else
            echo "Docker服务未运行，因此跳过编译docker镜像~"
        fi
    else
        echo "GOOS版本未识别~"
    fi

}

#清空指定目录下的内容（只执行一次）
mkdir -p ${OUTPUT_PATH}
echo ${OUTPUT_PATH}

#清理历史安装包
#rm -rf ./build/*
#echo "历史包清理完成..."

#循环编译指定版本的包
for GOOS in "${GOOS_LIST[@]}"; do
    for GOARCH in "${GOARCH_LIST[@]}"; do
        #遍历数组后拿到的<临时变量>拼接成<新的包名>
        baoName="gomessage-${VERSION}-${GOOS}-${GOARCH}"
        build
    done
done
