#!/usr/bin/env bash

#去掉不必要的包
go mod tidy

#编译之前：删除不必要的东西，不确定是否会被编译进去，但是以防万一比较好
rm -rf ./*.log
rm -rf ./*.tar.gz

#修改runmode参数，让编译后的应用程序运行在prod环境中
gsed -i '/^runmode/c runmode = prod' ./conf/app.conf

#交叉编译，在mac上编译linux上的部署包
bee pack -be GOOS=linux -be GOARCH=amd64 -exr 'build.sh' -v true

#编译完成之后再把runmode参数修改回来，这样下次都不用手动修改了
gsed -i '/^runmode/c runmode = dev' ./conf/app.conf
