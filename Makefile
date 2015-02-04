SHELL=/bin/bash

all: build

deps:
	go get -t -v ./...

test: build
	go test -v
	./cli-unit -v README.md *_test.md

build:
	go build


clean:
	rm -rf ./cli-unit
	rm -rf ./build

packages: build gzip

gzip: deps golang-crosscompile golang-buildsetup
	source golang-crosscompile/crosscompile.bash; \
	mkdir -p build/output; \
	go-darwin-386 build -o cli-unit; \
	gzip -c cli-unit > build/output/cli-unit-Darwin-386.gz; \
	go-darwin-amd64 build -o cli-unit; \
	gzip -c cli-unit > build/output/cli-unit-Darwin-x86_64.gz; \
	go-linux-386 build -o cli-unit; \
	gzip -c cli-unit > build/output/cli-unit-Linux-386.gz; \
	go-linux-amd64 build -o cli-unit; \
	gzip -c cli-unit > build/output/cli-unit-Linux-x86_64.gz

golang-buildsetup: golang-crosscompile
	source golang-crosscompile/crosscompile.bash; \
	go-crosscompile-build darwin/386; \
	go-crosscompile-build darwin/amd64; \
	go-crosscompile-build linux/386; \
	go-crosscompile-build linux/amd64

golang-crosscompile:
	git clone https://github.com/davecheney/golang-crosscompile.git

