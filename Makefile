.DEFAULT_GOAL := build

build:
	@echo "Building the project..."
	go build -buildmode=archive -o mylib.a .