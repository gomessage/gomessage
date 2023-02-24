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
all:  clean  start  swagger  build_linux  build_mac  build_windows  end


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
# Target：编译为Linux发行版（实际封装到容器里的内容）
######################################
.PHONY: build_linux
build_linux: packagePath:=${NAME}-${VERSION}-linux-x64
build_linux:
	mkdir -p "${OUTPUT_PATH}/${packagePath}"
	GOARCH=amd64 \
	GOOS=linux \
	CGO_ENABLED=${CGO_STATUS} \
	go build -ldflags '-extldflags "-static"' -o "${OUTPUT_PATH}/${packagePath}/${NAME}" ./main.go
	cp -rf ./config "${OUTPUT_PATH}/${packagePath}/"
	cp -rf ./assets "${OUTPUT_PATH}/${packagePath}/"
	tar -zcvf "${OUTPUT_PATH}/${packagePath}.tar.gz" -C ${OUTPUT_PATH} ${packagePath}
	ls -alh "${OUTPUT_PATH}/${packagePath}/"


######################################
# Target：编译为Mac发行版（本地调试使用）
######################################
.PHONY: build_mac
build_mac: packagePath:=${NAME}-${VERSION}-mac-x64
build_mac:
	mkdir -p "${OUTPUT_PATH}/${packagePath}"
	GOARCH=amd64 \
	GOOS=darwin \
	CGO_ENABLED=${CGO_STATUS} \
	go build -o "${OUTPUT_PATH}/${packagePath}/${NAME}" ./main.go
	cp -rf ./config "${OUTPUT_PATH}/${packagePath}/"
	cp -rf ./assets "${OUTPUT_PATH}/${packagePath}/"
	tar -zcvf "${OUTPUT_PATH}/${packagePath}.tar.gz" -C ${OUTPUT_PATH} ${packagePath}
	ls -alh "${OUTPUT_PATH}/${packagePath}/"


######################################
# Target：编译为Windows发行版
######################################
.PHONY: build_windows
build_windows: packagePath:=${NAME}-${VERSION}-windows-x64
build_windows:
	mkdir -p "${OUTPUT_PATH}/${packagePath}"
	GOARCH=amd64 \
	GOOS=windows \
	CGO_ENABLED=${CGO_STATUS} \
	go build -o "${OUTPUT_PATH}/${packagePath}/${NAME}.exe" ./main.go
	cp -rf ./config "${OUTPUT_PATH}/${packagePath}/"
	cp -rf ./assets "${OUTPUT_PATH}/${packagePath}/"
	tar -zcvf "${OUTPUT_PATH}/${packagePath}.tar.gz" -C ${OUTPUT_PATH} ${packagePath}
	ls -alh "${OUTPUT_PATH}/${packagePath}/"


######################################
# Target：结束之前做些什么
######################################
.PHONY: end
end:
	ls -alh ${OUTPUT_PATH}


######################################
# Target：编译为docker镜像
######################################
.PHONY: docker
docker:
	docker build -t ${NAME}:latest -f Dockerfile .
	docker login harbor.zmlearn.com -u taycc -p xxxxxx
	docker tag ${NAME}:latest harbor.zmlearn.com/triton/triton:${DATE_NOW}_${OFFICE_ENV}
	docker push harbor.zmlearn.com/triton/triton:${DATE_NOW}_${OFFICE_ENV}
