VERSION := 1.0.9
OUTPUT_PATH := ./build/${VERSION}/
EXRS := '^build|^bak|^docs|^swagger|^deploy|^LICENSE|.sh$$|.sqlite3$$|.bak$$|.md$$|prof$$'

build-mac: baoName := gomessage-${VERSION}-mac-amd64
build-linux: baoName := gomessage-${VERSION}-linux-amd64
build-windows: baoName := gomessage-${VERSION}-windows-amd64





.PHONY: clean start build-linux build-windows build-mac end
all: clean start build-linux build-windows build-mac end

clean:
	rm -rf ./build/*
	rm -rf ./*.log
	rm -rf ./*.tar.gz
	rm -rf ./GoMessage
	rm -rf ./go_build_GoMessage

start:
	go mod tidy
	gsed -i '/^runmode/c runmode = prod' ./conf/app.conf

end:
	gsed -i '/^runmode/c runmode = dev' ./conf/app.conf

build-linux:
	mkdir -p ${OUTPUT_PATH}${baoName}/

	GOARCH=amd64 GOOS=linux CGO_ENABLED=1 \
	CGO_LDFLAGS="-static" \
	CC=x86_64-linux-musl-gcc \
	CXX=x86_64-linux-musl-g++ \
	bee pack -a gomessage -o "${OUTPUT_PATH}${baoName}/" -exr ${EXRS}

	cp ./deploy/install.sh ${OUTPUT_PATH}${baoName}/
	cp ./deploy/uninstall.sh ${OUTPUT_PATH}${baoName}/

	tar -zcvf ${OUTPUT_PATH}${baoName}.tar.gz -C ${OUTPUT_PATH} ${baoName}

build-windows:
	mkdir -p ${OUTPUT_PATH}${baoName}/

	GOARCH=amd64 GOOS=windows CGO_ENABLED=1  \
	CGO_CFLAGS="-g -O2 -Wno-return-local-addr" \
	CC=x86_64-w64-mingw32-gcc \
	CXX=x86_64-w64-mingw32-g++ \
	bee pack -a gomessage -o "${OUTPUT_PATH}${baoName}/" -exr ${EXRS}

	tar -zcvf ${OUTPUT_PATH}${baoName}.tar.gz -C ${OUTPUT_PATH} ${baoName}

build-mac:
	mkdir -p ${OUTPUT_PATH}${baoName}/

	GOARCH=amd64 GOOS=darwin CGO_ENABLED=1 \
	bee pack -a gomessage -o "${OUTPUT_PATH}${baoName}/" -exr ${EXRS}

	tar -zcvf ${OUTPUT_PATH}${baoName}.tar.gz -C ${OUTPUT_PATH} ${baoName}
