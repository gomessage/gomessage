OUTPUT_PATH := ./build/
EXRS := '^build|^bak|^docs|^swagger|^deploy|^LICENSE|^goo|^Makefile|^imageSync|^Dockerfile|.sh$$|.sqlite3$$|.bak$$|.md$$|prof$$'

build-mac: baoName := gomessage-${VERSION}-mac-amd64
build-linux: baoName := gomessage-${VERSION}-linux-amd64
build-windows: baoName := gomessage-${VERSION}-windows-amd64





.PHONY: clean start build-linux build-linux-mac end
all: clean start build-linux end


clean:
	rm -rf ./build/*
	rm -rf ./*.log
	rm -rf ./*.tar.gz
	rm -rf ./GoMessage
	rm -rf ./go_build_GoMessage


start:
	go mod tidy
	sed -i '/^runmode/c runmode = prod' ./conf/app.conf


end:
	sed -i '/^runmode/c runmode = dev' ./conf/app.conf


build-linux:
	mkdir -p ${OUTPUT_PATH}
	GOARCH=amd64 \
	GOOS=linux \
	CGO_ENABLED=1 \
	bee pack -a gomessage -o "${OUTPUT_PATH}" -exr ${EXRS}
	docker build -t cicd:gomessage -f Dockerfile .
    imageSync -i cicd:gomessage

build-linux-mac:
	mkdir -p ${OUTPUT_PATH}
	GOARCH=amd64 \
	GOOS=linux \
	CGO_ENABLED=1 \
	CGO_LDFLAGS="-static" \
	CC=x86_64-linux-musl-gcc \
	CXX=x86_64-linux-musl-g++ \
	bee pack -a gomessage -o "${OUTPUT_PATH}" -exr ${EXRS}
