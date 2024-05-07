.DEFAULT_GOAL := build

BINNARY_NAME := ch8Program

.PHONY: build run clean

build:
	go build -o $(BINNARY_NAME) main.go

run: build
	./$(BINNARY_NAME)

clean:
	go clean
	rm -f $(BINNARY_NAME)