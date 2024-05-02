.DEFAULT_GOAL := build

BINARY_NAME := ch5program

.PHONY: build run clean

build:
	go build -o $(BINARY_NAME) main.go

run: build
	./$(BINARY_NAME)

clean:
	go clean -i