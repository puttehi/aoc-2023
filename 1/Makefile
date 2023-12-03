SHELL:=/bin/bash

all: build run

test:
	go test -v ./src/...

clean:
	rm -r build/*

build: test
	go build -o build/day-1 ./src/...

run:
	./build/day-1
