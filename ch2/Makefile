.DEFAULT_GOAL := build

BINARY_NAME := ch2program

.PHONY: build run

build:
	go build -o $(BINARY_NAME) main.go

run: build
	./$(BINARY_NAME)

clean:
	go clean -i