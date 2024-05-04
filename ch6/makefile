.DEFAULT_TARGET := build

BINARY_NAME := ch6program

.PHONY: build run clean

build:
	go build -o $(BINARY_NAME) -gcflags="-m" main.go

run: build
	GODEBUG=gctrace=1 go run main.go

clean:
	go clean -i

