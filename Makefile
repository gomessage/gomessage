# 编译新版本时只需修改<VERSION>的变量值即可（格式：x.x.x）
#
# 常用命令：
#
#	make all  --->  同时编译linux、mac、windows三个环境的tar包
#
#	make docker --->  编译docker镜像
#
#	make swagger  --->  生成swagger文档



######################################
# 全局变量
######################################
#要编译的命令名称
NAME := gomessage
#版本
VERSION := 2.0.1
#编译输出目录
OUTPUT_PATH := ./build/${VERSION}
#是否开启cgo（0代表不开启，1代表开启）
CGO_STATUS := 0
#当前时间
DATE_NOW := $(shell date "+%Y%m%d_%H%M%S")


######################################
# 指定缺省状态下执行哪些Target
######################################
all:  clean  start  swagger  build_mac  build_windows  build_linux  end


######################################
# Target：清理开发目录
######################################
.PHONY: clean
clean:
	mkdir -p "${OUTPUT_PATH}"
	echo "编译输出目录为：${OUTPUT_PATH}"
	rm -rf ./tmp
	rm -rf ./build/*
	rm -rf ./*.log
	rm -rf ./*.db
	rm -rf ./*.tar.gz


######################################
# Target：处理依赖
######################################
.PHONY: start
start:
	go mod tidy


######################################
# Target：生成swagger文件
######################################
.PHONY: swagger
swagger:
	go install github.com/swaggo/swag/cmd/swag@latest
	mkdir -p assets
	swag init -o assets/docs


######################################
# Target：编译为Mac发行版（本地调试使用）
######################################
.PHONY: build_mac
build_mac: packageName:=${NAME}-${VERSION}-mac-x64
build_mac:
	mkdir -p "${OUTPUT_PATH}/${packageName}"
	GOARCH=amd64 \
	GOOS=darwin \
	CGO_ENABLED=${CGO_STATUS} \
	go build -o "${OUTPUT_PATH}/${packageName}/${NAME}" ./main.go
	cp -rf ./config "${OUTPUT_PATH}/${packageName}/"
	cp -rf ./assets "${OUTPUT_PATH}/${packageName}/"
	tar -zcvf "${OUTPUT_PATH}/${packageName}.tar.gz" -C ${OUTPUT_PATH} ${packageName}
	ls -alh "${OUTPUT_PATH}/${packageName}/"


######################################
# Target：编译为Windows发行版
######################################
.PHONY: build_windows
build_windows: packageName:=${NAME}-${VERSION}-windows-x64
build_windows:
	mkdir -p "${OUTPUT_PATH}/${packageName}"
	GOARCH=amd64 \
	GOOS=windows \
	CGO_ENABLED=${CGO_STATUS} \
	go build -o "${OUTPUT_PATH}/${packageName}/${NAME}.exe" ./main.go
	cp -rf ./config "${OUTPUT_PATH}/${packageName}/"
	cp -rf ./assets "${OUTPUT_PATH}/${packageName}/"
	tar -zcvf "${OUTPUT_PATH}/${packageName}.tar.gz" -C ${OUTPUT_PATH} ${packageName}
	ls -alh "${OUTPUT_PATH}/${packageName}/"


######################################
# Target：编译为Linux发行版（实际封装到容器里的内容）
######################################
.PHONY: build_linux
build_linux: packageName:=${NAME}-${VERSION}-linux-x64
build_linux:
	mkdir -p "${OUTPUT_PATH}/${packageName}"
	GOARCH=amd64 \
	GOOS=linux \
	CGO_ENABLED=${CGO_STATUS} \
	go build -ldflags '-extldflags "-static"' -o "${OUTPUT_PATH}/${packageName}/${NAME}" ./main.go
	cp -rf ./config "${OUTPUT_PATH}/${packageName}/"
	cp -rf ./assets "${OUTPUT_PATH}/${packageName}/"
	tar -zcvf "${OUTPUT_PATH}/${packageName}.tar.gz" -C ${OUTPUT_PATH} ${packageName}
	@echo "\n---------编译封装image专用的tar包：开始---------\n"
	tar -zcvf "${OUTPUT_PATH}/${NAME}.tar.gz" -C ${OUTPUT_PATH} ${packageName}
	@echo "\n---------编译封装image专用的tar包：完成---------\n"
	ls -alh "${OUTPUT_PATH}/${packageName}/"


######################################
# Target：编译为docker镜像
######################################
#DOCKER_IS_RUN := "-"
.PHONY: docker
docker: DOCKER_SCAN_SUGGEST := False
docker:
	@echo "\n---------版本latest---------\n"
	docker build -t gomessage/gomessage:latest -f ./deploy/Dockerfile  "${OUTPUT_PATH}/"
	@echo "\n---------开始制作镜像，版本${VERSION}---------\n"
	docker build -t gomessage/gomessage:${VERSION} -f ./deploy/Dockerfile  "${OUTPUT_PATH}/"
	@echo "\n---------镜像制作完成，版本${VERSION}---------\n"


######################################
# Target：结束之前做些什么
######################################
.PHONY: end
end:
	ls -alh ${OUTPUT_PATH}