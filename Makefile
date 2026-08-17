# 编译新版本时只需修改<VERSION>的变量值即可（格式：x.x.x）
#
# 版本管理约定：
#	VERSION 是全局唯一版本号来源；Chart.yaml 的 version/appVersion 由构建流程自动同步
#	正式发布前，HEAD 上必须已打好 v$(VERSION) 的 tag（make check 会自动校验）
#
# 常用命令：
#
#	make check  --->  发布前置校验（工作区干净 + tag 存在）
#
#	make all  --->  同时编译linux、mac、windows三个环境的tar包
#
#	make docker --->  编译docker镜像
#
#	make docker_push
#
#	make package_push
#
#	make release  --->  一键发布：check + 编译 + 多架构镜像 + Helm + 上传GitHub Release
#
#	make swagger
#




######################################
# 全局变量
######################################
#要编译的命令名称
NAME := gomessage
#版本（唯一版本号来源，发布其他版本时只需改这里）
VERSION := 3.0.0
#对应的git tag
GIT_TAG := v$(VERSION)
#编译输出目录
OUTPUT_PATH := ./build/${VERSION}
#是否开启cgo（0代表不开启，1代表开启）
CGO_STATUS := 1
#当前时间
DATE_NOW := $(shell date "+%Y%m%d_%H%M%S")
#前端构建标记（避免多平台编译时重复执行npm build）
FRONTEND_STAMP := ./build/.frontend_done


######################################
# 指定缺省状态下执行哪些Target
######################################
all: clean start swagger build_frontend build_mac_arm64 build_windows build_linux build_linux_arm64 end


######################################
# Target：发布前置校验
######################################
.PHONY: check
check:
	@echo "\n---------发布前置校验---------\n"
	@git diff-index --quiet HEAD -- || { echo "ERROR: 工作区存在未提交变更，请先提交或stash"; exit 1; }
	@git tag --points-at HEAD | grep -qx '$(GIT_TAG)' || { echo "ERROR: 当前HEAD上不存在tag $(GIT_TAG)，请先执行: git tag -a $(GIT_TAG) -m \"release: $(GIT_TAG)\""; exit 1; }
	@echo "校验通过：工作区干净，HEAD已标记为 $(GIT_TAG)\n"


######################################
# Target：清理开发目录
######################################
.PHONY: clean
clean:
	rm -rf ./tmp
	rm -rf ./build/*
	rm -rf ./*.log
	rm -rf ./*.db
	rm -rf ./*.tar.gz
	rm -rf ./config/*.db
	rm -rf ./*.tgz
	mkdir -p "${OUTPUT_PATH}"
	echo "编译输出目录为：${OUTPUT_PATH}"


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
	go install github.com/swaggo/swag/cmd/swag@v1.8.12
	mkdir -p assets
	swag init -o assets/docs


######################################
# Target：构建前端静态资源（幂等，一次make流程内只真正构建一次）
######################################
.PHONY: build_frontend
build_frontend:
	@if [ -f "$(FRONTEND_STAMP)" ]; then \
		echo "前端已构建过，跳过（如需强制重建请先 make clean）"; \
	else \
		echo "\n---------开始构建前端---------\n"; \
		cd vue && NODE_OPTIONS=--openssl-legacy-provider npm run build; \
		touch "$(FRONTEND_STAMP)"; \
		echo "\n---------前端构建完成---------\n"; \
	fi


######################################
# Target：编译为Mac的x86_64发行版（本地调试使用）
######################################
.PHONY: build_mac
build_mac: build_frontend
build_mac: packageName:=${NAME}-${VERSION}-mac-amd64
build_mac:
	mkdir -p "${OUTPUT_PATH}/${packageName}"
	GOARCH=amd64 \
	GOOS=darwin \
	CGO_ENABLED=${CGO_STATUS} \
	go build -ldflags='-s -w' -o "${OUTPUT_PATH}/${packageName}/${NAME}" ./main.go
	cp -rf ./config "${OUTPUT_PATH}/${packageName}/"
	cp -rf ./assets "${OUTPUT_PATH}/${packageName}/"
	tar -zcvf "${OUTPUT_PATH}/${packageName}.tar.gz" -C ${OUTPUT_PATH} ${packageName}
	ls -alh "${OUTPUT_PATH}/${packageName}/"


######################################
# Target：编译为Mac的arm发行版（本地调试使用）
######################################
.PHONY: build_mac_arm64
build_mac_arm64: build_frontend
build_mac_arm64: packageName:=${NAME}-${VERSION}-mac-arm64
build_mac_arm64:
	mkdir -p "${OUTPUT_PATH}/${packageName}"
	GOARCH=arm64 \
	GOOS=darwin \
	CGO_ENABLED=${CGO_STATUS} \
	go build -ldflags='-s -w' -o "${OUTPUT_PATH}/${packageName}/${NAME}" ./main.go
	cp -rf ./config "${OUTPUT_PATH}/${packageName}/"
	cp -rf ./assets "${OUTPUT_PATH}/${packageName}/"
	tar -zcvf "${OUTPUT_PATH}/${packageName}.tar.gz" -C ${OUTPUT_PATH} ${packageName}
	ls -alh "${OUTPUT_PATH}/${packageName}/"


######################################
# Target：编译为Windows发行版
######################################
.PHONY: build_windows
build_windows: build_frontend
build_windows: packageName:=${NAME}-${VERSION}-windows-amd64
build_windows:
	mkdir -p "${OUTPUT_PATH}/${packageName}"
	GOARCH=amd64 \
	GOOS=windows \
	CGO_CFLAGS="-g -O2 -Wno-return-local-addr" \
	CC=x86_64-w64-mingw32-gcc \
	CXX=x86_64-w64-mingw32-g++ \
	CGO_ENABLED=${CGO_STATUS} \
	go build -ldflags='-s -w -extldflags "-static"' -o "${OUTPUT_PATH}/${packageName}/${NAME}.exe" ./main.go
	cp -rf ./config "${OUTPUT_PATH}/${packageName}/"
	cp -rf ./assets "${OUTPUT_PATH}/${packageName}/"
	tar -zcvf "${OUTPUT_PATH}/${packageName}.tar.gz" -C ${OUTPUT_PATH} ${packageName}
	ls -alh "${OUTPUT_PATH}/${packageName}/"


######################################
# Target：编译为Linux发行版amd64（实际封装到容器里的内容）
######################################
.PHONY: build_linux
build_linux: build_frontend
build_linux: packageName:=${NAME}-${VERSION}-linux-amd64
build_linux:
	mkdir -p "${OUTPUT_PATH}/${packageName}"
	GOARCH=amd64 \
	GOOS=linux \
	CGO_LDFLAGS="-static" \
	CC=x86_64-linux-musl-gcc \
	CXX=x86_64-linux-musl-g++ \
	CGO_ENABLED=${CGO_STATUS} \
	go build -ldflags='-s -w -extldflags "-static"' -o "${OUTPUT_PATH}/${packageName}/${NAME}" ./main.go
	cp -rf ./config "${OUTPUT_PATH}/${packageName}/"
	cp -rf ./assets "${OUTPUT_PATH}/${packageName}/"
	tar -zcvf "${OUTPUT_PATH}/${packageName}.tar.gz" -C ${OUTPUT_PATH} ${packageName}
	ls -alh "${OUTPUT_PATH}/${packageName}/"


######################################
# Target：编译为Linux发行版arm64（用于arm64架构的容器/服务器）
######################################
.PHONY: build_linux_arm64
build_linux_arm64: build_frontend
build_linux_arm64: packageName:=${NAME}-${VERSION}-linux-arm64
build_linux_arm64:
	mkdir -p "${OUTPUT_PATH}/${packageName}"
	GOARCH=arm64 \
	GOOS=linux \
	CGO_LDFLAGS="-static" \
	CC=aarch64-linux-musl-gcc \
	CXX=aarch64-linux-musl-g++ \
	CGO_ENABLED=${CGO_STATUS} \
	go build -ldflags='-s -w -extldflags "-static"' -o "${OUTPUT_PATH}/${packageName}/${NAME}" ./main.go
	cp -rf ./config "${OUTPUT_PATH}/${packageName}/"
	cp -rf ./assets "${OUTPUT_PATH}/${packageName}/"
	tar -zcvf "${OUTPUT_PATH}/${packageName}.tar.gz" -C ${OUTPUT_PATH} ${packageName}
	ls -alh "${OUTPUT_PATH}/${packageName}/"


######################################
# Target：结束之前做些什么
######################################
.PHONY: end
end:
	ls -alh ${OUTPUT_PATH}


######################################
# Target：编译docker镜像
######################################
.PHONY: docker
docker: build_linux
docker: packageName := ${NAME}-${VERSION}-linux-amd64
docker:
	@echo "\n---------版本latest---------\n"
	@docker build -t gomessage/gomessage:latest -f ./docker/Dockerfile  "${OUTPUT_PATH}/${packageName}"
	@echo "\n---------开始制作镜像，版本${VERSION}---------\n"
	@docker build -t gomessage/gomessage:${VERSION} -f ./docker/Dockerfile  "${OUTPUT_PATH}/${packageName}"
	@echo "\n---------镜像制作完成，版本${VERSION}---------\n"


######################################
# Target：推送docker多架构镜像及Helm Chart
######################################
.PHONY: docker_push
docker_push: build_linux build_linux_arm64
docker_push:
	# 注意：需要 docker buildx、gsed、helm 等工具支持
	#docker login --username=$(DOCKER_HUB_USERNAME)
	#docker buildx rm mybuilder
	#docker buildx create --name mybuilder --bootstrap --use
	@echo "\n---------开始制作多架构镜像，版本${VERSION}---------\n"
	@docker buildx build --platform linux/amd64 -t gomessage/gomessage:${VERSION} -t gomessage/gomessage:latest -f ./docker/Dockerfile "${OUTPUT_PATH}/${NAME}-${VERSION}-linux-amd64" --push
	@docker buildx build --platform linux/arm64 -t gomessage/gomessage:${VERSION} -t gomessage/gomessage:latest -f ./docker/Dockerfile "${OUTPUT_PATH}/${NAME}-${VERSION}-linux-arm64" --push
	@echo "\n---------镜像制作完成，版本${VERSION}---------\n"
	@echo
	@gsed -i '/version:/c version: ${VERSION}' ./docker/helm/Chart.yaml
	@gsed -i '/appVersion:/c appVersion: ${VERSION}' ./docker/helm/Chart.yaml
	helm package ./docker/helm
	helm coding-push gomessage-${VERSION}.tgz gomessage
	rm -rf ./*.tgz
	@echo "\n---------制作Helm Chart完成，版本${VERSION}---------\n"



######################################
# Target：推送Helm Chart（不重新构建镜像时使用）
######################################
.PHONY: helm_push
helm_push:
	# 注意：需要 gsed、helm coding-push 等工具支持
	@gsed -i '/version:/c version: ${VERSION}' ./docker/helm/Chart.yaml
	@gsed -i '/appVersion:/c appVersion: ${VERSION}' ./docker/helm/Chart.yaml
	helm package ./docker/helm
	helm coding-push gomessage-${VERSION}.tgz gomessage
	rm -rf ./*.tgz
	@echo "\n---------制作Helm Chart完成，版本${VERSION}---------\n"



######################################
# Target：推送package到github（创建Release并上传安装包）
######################################
.PHONY: package_push
package_push:
	@test -n "$$Github_Authorization" -a -n "$$Github_Token" || { echo "ERROR: 请先 export Github_Authorization=\"Authorization\" 与 Github_Token=\"Bearer <token>\""; exit 1; }
	@go run ./cmd/uploads --version=${VERSION}


######################################
# Target：一键发布（校验 + 编译 + 镜像 + Helm + GitHub Release）
######################################
.PHONY: release
release: check all docker_push package_push
	@echo "\n---------$(GIT_TAG) 发布完成---------\n"
