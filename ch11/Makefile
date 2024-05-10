.DEFAULT_GOAL := build

BINAY_NAME := myapp

.PHONY: build clean run

build:
	@echo "Building the project..."
	GOOS=linux GOARCH=amd64 go build -o ${BINAY_NAME} main.go

run: build
	@echo "Running the project..."
	./${BINAY_NAME} sinhala

clean:
	@echo "Cleaning up..."
	go clean
	rm -f ${BINAY_NAME}
