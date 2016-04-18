GOARCH ?= amd64
GOOS ?= linux

all : clean build

clean :
	rm -rf wrapper0

.wrapper0 :
	docker run --rm \
		-v ${CURDIR}:/usr/src/wrapper0 \
		-w /usr/src/wrapper0 \
		-e GOOS=${GOOS} \
		-e GOARCH=${GOARCH} \
		golang:1.6 \
		go build

build : .wrapper0
