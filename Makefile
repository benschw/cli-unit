SHELL=/bin/bash

all: build


test: build
	go test $(go list ./... | grep -v '/vendor/')
	./cli-unit -v README.md *_test.md

build:
	go build


clean:
	rm -rf ./cli-unit
	rm -rf ./build
	rm -rf ./.cli-unit

packages: build xcompile

xcompile:
	mkdir -p build/output
	env GOOS=darwin GOARCH=386 go build -o cli-unit
	gzip -c cli-unit > build/output/cli-unit-Darwin-386.gz
	env GOOS=darwin GOARCH=amd64 go build -o cli-unit
	gzip -c cli-unit > build/output/cli-unit-Darwin-x86_64.gz
	env GOOS=linux GOARCH=386 go build -o cli-unit
	gzip -c cli-unit > build/output/cli-unit-Linux-386.gz
	env GOOS=linux GOARCH=amd64 go build -o cli-unit
	gzip -c cli-unit > build/output/cli-unit-Linux-x86_64.gz

