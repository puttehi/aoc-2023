SHELL:=/bin/bash

.PHONY:$(MAKECMDGOALS)

all: clean test build run

run: build
	./build/4

build: test
	go build -o ./build/4 ./cmd/4/...

clean:
	rm -rf ./build/*

test:
	go test -v ./...
