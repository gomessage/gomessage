#!/usr/bin/env bash

go mod tidy

rm -rf ./*.log

rm -rf ./*.tar.gz

gsed -i  '/^runmode/c runmode = prod' conf/app.conf

#交叉编译，在mac上编译linux上的部署包
bee pack -be GOOS=linux -be GOARCH=amd64
