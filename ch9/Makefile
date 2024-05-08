.DEFAULT_GOAL := build

APPLICATION_NAME := main

.PHONY: build run

build:
	go build -o ${APPLICATION_NAME}

run: build
	./${APPLICATION_NAME}

clean:
	rm -f main