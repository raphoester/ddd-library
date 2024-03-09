BINARY_NAME=library

MAIN_PACKAGE=./cmd/api

WIRE=./vendor/github.com/google/wire/cmd/wire

wire:
	@echo "Generating dependency injection..."
	@go generate ./...

build: wire
	@echo "Building app..."
	@go build -o ${BINARY_NAME} ${MAIN_PACKAGE}

clean:
	@echo "Cleaning..."
	@go clean
	@rm -f ${BINARY_NAME}

run: wire
	@echo "Running app..."
	@go run ${MAIN_PACKAGE}


.PHONY: all wire build clean
