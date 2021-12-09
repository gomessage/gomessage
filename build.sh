#!/usr/bin/env bash

#交叉编译，在mac上编译linux上的部署包
bee pack -be GOOS=linux -be GOARCH=amd64