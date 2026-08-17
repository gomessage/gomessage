# 发布流程（三步走）：
#
#	make release  --->  锁版本：交互式输入新版本号，自动完成
#	                    合并dev→main、版本号辐射、提交、打tag、推送GitHub（仅限main分支）
#
#	make build    --->  编译：编译四平台安装包（mac arm64 / windows / linux amd64+arm64）
#
#	make publish  --->  推送：凭证校验后，推送多架构Docker镜像、上传GitHub Release
#
# 日常开发时在 dev 分支随意提交推送，互不影响。
# 发布新版本时只需在 release 交互提示中输入新版本号（如 3.0.1），无需手改任何文件。
#




######################################
# 全局变量
######################################
#要编译的命令名称
NAME := gomessage
#当前版本（由 make release 自动维护，请勿手改）
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
# Target：分支守卫（发布链路命令仅允许在 main 分支执行，防止在 dev 上误操作）
######################################
.PHONY: guard_main
guard_main:
	@BRANCH="$$(git rev-parse --abbrev-ref HEAD)" ; \
	[ "$${BRANCH}" = "main" ] || { echo "ERROR: 发布相关命令仅允许在 main 分支执行（当前分支：$${BRANCH}），请先切换到 main" ; exit 1; }


######################################
# 指定缺省状态下执行哪些Target
######################################
build: guard_main clean start swagger build_frontend build_mac_arm64 build_windows build_linux build_linux_arm64 end

# 兼容旧命令
all: build


######################################
# Target：查看发布流程说明
######################################
.PHONY: help
help:
	@echo "\n发布流程（三步走）："
	@echo "  make release  ->  锁版本：交互输入版本号，自动合并dev/辐射版本/提交/打tag/推送"
	@echo "  make build    ->  编译：四平台安装包（mac arm64 / windows / linux amd64+arm64）"
	@echo "  make publish  ->  推送：多架构Docker镜像 + GitHub Release安装包"
	@echo "  （以上命令仅限 main 分支执行，其它分支会被强制拦截）"
	@echo "\n辅助命令："
	@echo "  make status   ->  查看当前分支/版本/tag/产物状态"
	@echo "  make creds    ->  单独校验发布凭证\n"


######################################
# Target：查看当前发布状态
######################################
.PHONY: status
status:
	@echo "\n---------当前状态---------"
	@echo "分支：$$(git rev-parse --abbrev-ref HEAD)"
	@echo "版本：$(VERSION)"
	@if git tag --points-at HEAD | grep -qx '$(GIT_TAG)'; then echo "tag ：$(GIT_TAG) ✅ 已锁定在当前HEAD"; else echo "tag ：$(GIT_TAG) ⚠️ 未锁定在当前HEAD（需 make release）"; fi
	@if [ -d "${OUTPUT_PATH}" ]; then echo "产物：${OUTPUT_PATH} ✅ 已存在"; else echo "产物：${OUTPUT_PATH} ⚠️ 不存在（需 make build）"; fi
	@echo "--------------------------\n"


######################################
# Target：锁版本（交互式，只能在main分支执行）
# 自动完成：合并dev → 输入新版本号 → 版本号辐射 → 提交 → 打tag → 推送GitHub
######################################
.PHONY: release
release:
	@echo "\n---------开始锁定版本---------\n"
	@[ "$$(git rev-parse --abbrev-ref HEAD)" = "main" ] || { echo "ERROR: make release 只能在 main 分支执行"; exit 1; }
	@git update-index --refresh >/dev/null 2>&1; git diff-index --quiet HEAD -- || { echo "ERROR: 工作区存在未提交变更，请先提交或stash"; exit 1; }
	@git merge-base --is-ancestor dev main || { echo ">> dev 有未合并变更，自动合并到 main..."; git merge dev --no-edit; }
	@printf "当前版本：$(VERSION)\n请输入要发布的版本号（如 3.0.1 或 3.1.0，重锁可输入当前版本号）：" ; \
	read NEW_VERSION ; \
	echo "$$NEW_VERSION" | grep -Eq '^[0-9]+\.[0-9]+\.[0-9]+$$' || { echo "ERROR: 版本号格式必须是 x.y.z"; exit 1; } ; \
	if git tag -l | grep -qx "v$$NEW_VERSION"; then \
		printf "tag v$${NEW_VERSION} 已存在，是否重新锁定（tag将移动到当前HEAD）？[y/N] " ; \
	else \
		printf "确认发布 v$${NEW_VERSION}（当前 $(VERSION) → v$${NEW_VERSION}）？[y/N] " ; \
	fi ; \
	read CONFIRM ; \
	[ "$$CONFIRM" = "y" ] || { echo "已取消"; exit 1; } ; \
	echo ">> 版本号辐射：Makefile / Chart.yaml" ; \
	gsed -i -E "s/^VERSION := .*/VERSION := $$NEW_VERSION/" Makefile ; \
	gsed -i "/version:/c version: $$NEW_VERSION" ./docker/helm/Chart.yaml ; \
	gsed -i "/appVersion:/c appVersion: $$NEW_VERSION" ./docker/helm/Chart.yaml ; \
	echo ">> 提交并打tag" ; \
	git add Makefile ./docker/helm/Chart.yaml ; \
	if git diff --cached --quiet; then echo ">> 版本号无变化，跳过提交"; else git commit -m "release: v$$NEW_VERSION"; fi ; \
	git tag -af "v$$NEW_VERSION" -m "release: v$$NEW_VERSION" ; \
	echo ">> 推送 main 和 tag 到远程（重锁时tag会强制更新）" ; \
	git push github main ; \
	git push -f github "refs/tags/v$$NEW_VERSION" ; \
	echo "\n---------版本已锁定：v$$NEW_VERSION---------" ; \
	echo ">> 下一步执行 make build 和 make publish 开始编译和分发\n"


######################################
# Target：发布前置校验（只验tag在HEAD上，不验工作区干净：
# make build 会重新生成swagger/前端产物，产生正常diff，若验干净会与tag校验形成死锁）
######################################
.PHONY: check
check:
	@echo "\n---------发布前置校验---------\n"
	@git tag --points-at HEAD | grep -qx '$(GIT_TAG)' || { echo "ERROR: 当前HEAD上不存在tag $(GIT_TAG)，请先执行 make release 锁定版本"; exit 1; }
	@echo "校验通过：HEAD已标记为 $(GIT_TAG)\n"


######################################
# Target：发布凭证校验
######################################
.PHONY: creds
creds:
	@echo "\n---------凭证校验---------\n"
	@docker info >/dev/null 2>&1 || { echo "ERROR: Docker 未启动，请先启动 Docker Desktop"; exit 1; }
	@grep -q 'docker.io' ~/.docker/config.json 2>/dev/null || { echo "ERROR: 未登录 Docker Hub，请先执行 docker login"; exit 1; }
	@{ [ -n "$$Github_Token" ] || { command -v gh >/dev/null 2>&1 && gh auth token >/dev/null 2>&1; }; } || { echo "ERROR: 未检测到GitHub凭证：请 export Github_Token=\"Bearer <token>\"，或执行 gh auth login"; exit 1; }
	@echo "凭证校验通过\n"


######################################
# Target：推送（一键：凭证校验 + 推镜像 + 推Helm + 传GitHub Release）
# 前置条件：已执行过 make build
######################################
.PHONY: publish
publish: guard_main creds check docker_push package_push
	@echo "\n---------$(GIT_TAG) 发布完成---------\n"


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
		touch "$(CURDIR)/$(FRONTEND_STAMP)"; \
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
# Target：编译docker镜像（本地调试）
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
docker_push: guard_main
	# 注意：需要 docker buildx、gsed、helm 等工具支持；请先执行 make build
	@for p in linux-amd64 linux-arm64; do \
		[ -d "${OUTPUT_PATH}/${NAME}-${VERSION}-$${p}" ] || { echo "ERROR: 缺少编译产物 ${NAME}-${VERSION}-$${p}，请先执行 make build"; exit 1; } ; \
	done
	#docker login --username=$(DOCKER_HUB_USERNAME)
	#docker buildx rm mybuilder
	#docker buildx create --name mybuilder --bootstrap --use
	@echo "\n---------开始制作多架构镜像，版本${VERSION}---------\n"
	@docker buildx build --platform linux/amd64 -t gomessage/gomessage:${VERSION} -t gomessage/gomessage:latest -f ./docker/Dockerfile "${OUTPUT_PATH}/${NAME}-${VERSION}-linux-amd64" --push
	@docker buildx build --platform linux/arm64 -t gomessage/gomessage:${VERSION} -t gomessage/gomessage:latest -f ./docker/Dockerfile "${OUTPUT_PATH}/${NAME}-${VERSION}-linux-arm64" --push
	@echo "\n---------镜像制作完成，版本${VERSION}---------\n"



######################################
# Target：推送package到github（创建Release并上传安装包）
######################################
.PHONY: package_push
package_push: guard_main
	# 优先使用 Github_Token 环境变量；未设置时自动复用 gh CLI 登录态（gh auth token）
	@TOKEN="$$Github_Token" ; \
	if [ -z "$$TOKEN" ] && command -v gh >/dev/null 2>&1; then \
		GH_TOKEN=$$(gh auth token 2>/dev/null) ; \
		[ -n "$$GH_TOKEN" ] && TOKEN="Bearer $$GH_TOKEN" ; \
	fi ; \
	[ -n "$$TOKEN" ] || { echo "ERROR: 未检测到 GitHub凭证：请 export Github_Token=\"Bearer <token>\"，或执行 gh auth login"; exit 1; } ; \
	Github_Authorization="$${Github_Authorization:-Authorization}" Github_Token="$$TOKEN" go run ./cmd/uploads --version=${VERSION}
