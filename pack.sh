#!/usr/bin/env bash

VERSION="v1.0.0"
OUTPUT_PATH="./build/${VERSION}/"
GOOS_LIST=(darwin linux windows)
GOARCH_LIST=(amd64)

build() {
    #去掉不必要的包
    go mod tidy

    #编译之前：删除不必要的东西，不确定是否会被编译进去，但是以防万一比较好
    rm -rf ./*.log
    rm -rf ./*.tar.gz
    rm -rf ./GoMessage

    #修改runmode参数，让编译后的应用程序运行在prod环境中
    gsed -i '/^runmode/c runmode = prod' ./conf/app.conf

    #交叉编译，在mac上编译linux上的部署包
    bee pack -a "${baoName}" \
        -o "${OUTPUT_PATH}" \
        -exr "^pack.sh|^build|README.md|.sh\$|bak" \
        -be GOOS="${GOOS}" \
        -be GOARCH="${GOARCH}"

    #编译完成之后再把runmode参数修改回来，这样下次都不用手动修改了
    gsed -i '/^runmode/c runmode = dev' ./conf/app.conf

}

#清空指定目录下的内容（只执行一次）
echo ${OUTPUT_PATH}
rm -rf "${OUTPUT_PATH}"*
echo "历史包清理完成..."

#循环编译指定版本的包
for GOOS in "${GOOS_LIST[@]}"; do
    for GOARCH in "${GOARCH_LIST[@]}"; do
        baoName="GoMessage-${VERSION}-${GOOS}-${GOARCH}"
        build
    done

done
