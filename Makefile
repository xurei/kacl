.PHONY: setup test deps

mkfile_path 	:= $(abspath $(lastword $(MAKEFILE_LIST)))
mkfile_dir 		:= $(shell dirname $(mkfile_path))
#PROJECT_PATH 	:= $(shell cd ${mkfile_dir}/../.. && pwd)
PROJECT_PATH 	:= $(shell cd ${mkfile_dir} && pwd)
PROJECT_DIR 	:= $(shell basename ${PROJECT_PATH})
OUT_DIR 		:= ${PROJECT_PATH}/target
GO      		:= go
GO_SRC      	:= ${PROJECT_PATH}/src/main/golang

install:
	go install

deps:
	${GO} mod vendor
	#go get -v ./...
	${GO} get golang.org/x/lint/golint
#	${GO} get github.com/stretchr/testify

test:
	go test -v ${GO_SRC}/...
	go test -cover ${GO_SRC}/...
	${GOPATH}/bin/golint ${GO_SRC}/...
	go vet -all ${GO_SRC}
	#${GOPATH}/bin/gocyclo -over 10 .

clean:
	rm kacl-*.tar.gz

build: ARGS=-v
build: GOOS ?= linux
build: GOARCH ?= amd64
build: VERSION ?= latest
build:
	GOOS=${GOOS} GOARCH=${GOARCH} ${GO} build -o kacl ${GO_SRC}/main.go

releases: TAG=$(shell git describe --tags $(git rev-list --tags --max-count=1))
releases:
	GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -o kacl ${GO_SRC}/main.go    && tar czvf kacl-$(TAG)-linux_amd64.tar.gz kacl
	GOARCH=386   GOOS=linux CGO_ENABLED=0 go build -o kacl ${GO_SRC}/main.go    && tar czvf kacl-$(TAG)-linux_386.tar.gz kacl
	GOARCH=amd64 GOOS=darwin CGO_ENABLED=0 go build -o kacl ${GO_SRC}/main.go   && tar czvf kacl-$(TAG)-darwin_amd64.tar.gz kacl
	GOARCH=386   GOOS=darwin CGO_ENABLED=0 go build -o kacl ${GO_SRC}/main.go   && tar czvf kacl-$(TAG)-darwin_386.tar.gz kacl
	GOARCH=amd64 GOOS=windows CGO_ENABLED=0 go build -o kacl ${GO_SRC}/main.go  && tar czvf kacl-$(TAG)-windows_amd64.tar.gz kacl
	GOARCH=386   GOOS=windows CGO_ENABLED=0 go build -o kacl ${GO_SRC}/main.go  && tar czvf kacl-$(TAG)-windows_386.tar.gz kacl
	rm ./kacl
