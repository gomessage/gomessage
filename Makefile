OUTPUT_PATH := ./build
OFFICE_ENV := master
CMD_NAME := gomessage
DATE_NOW := $(shell date "+%Y%m%d_%H%M%S")

all: clean start  build_mac end


#清理开发目录
.PHONY: clean
clean:
	rm -rf ./build/*
	rm -rf ./*.log
	rm -rf ./*.tar.gz


#整理开发依赖
.PHONY: start
start:
	go mod tidy


#编译为Linux发行版（实际封装到容器中的内容）
.PHONY: build_linux
build_linux:
	mkdir -p "${OUTPUT_PATH}/linux"
	GOARCH=amd64 \
	GOOS=linux \
	CGO_ENABLED=1 \
	go build -ldflags '-extldflags "-static"' -o "${OUTPUT_PATH}/linux/${CMD_NAME}" ./main.go
	cp -rf ./config "${OUTPUT_PATH}/linux/"
	cp -rf ./assets "${OUTPUT_PATH}/linux/"


#编译为Mac发行版（本地调试使用）
.PHONY: build_mac
build_mac:
	mkdir -p "${OUTPUT_PATH}/mac"
	GOARCH=amd64 \
	GOOS=darwin \
	CGO_ENABLED=1 \
	go build -o "${OUTPUT_PATH}/mac/${CMD_NAME}" ./main.go
	cp -rf ./config "${OUTPUT_PATH}/mac/"
	cp -rf ./assets "${OUTPUT_PATH}/mac/"


#结束之前做些什么
.PHONY: end
end:
	ls -alh ${OUTPUT_PATH}


#编译为docker image（注意这里的上下文设定，Dockerfile中的基准上下文就是这里指定的）
.PHONY: docker
docker:
	docker build -t ${CMD_NAME}:latest -f Dockerfile .
	docker login harbor.zmlearn.com -u taycc -p xxxxxx
	docker tag ${CMD_NAME}:latest harbor.zmlearn.com/triton/triton:${DATE_NOW}_${OFFICE_ENV}
	docker push harbor.zmlearn.com/triton/triton:${DATE_NOW}_${OFFICE_ENV}

#这个模块如果有需要，需要手动来执行命令：make git
.PHONY: git
git:
	git add .
	git commit -am "docker_image_${DATE_NOW}_${OFFICE_ENV}"
	git tag "docker_image_${DATE_NOW}_${OFFICE_ENV}"
	git push --tags

