BINARY_NAME=library

MAIN_PACKAGE=./cmd/api

WIRE=./vendor/github.com/google/wire/cmd/wire

BCS := authentication borrow
PROTO_DIR := ./proto
CONTEXTS_DIR := ./internal/contexts

PROTOC := protoc
PROTOC_GEN_GO := protoc-gen-go
PROTOC_GEN_GO_GRPC := protoc-gen-go-grpc
PROTOC_GO_INJECT_TAG := protoc-go-inject-tag

wire:
	@echo "Generating dependency injection..."
	@go generate ./...

build: proto wire
	@echo "Building app..."
	@go build -o ${BINARY_NAME} ${MAIN_PACKAGE}

clean:
	@echo "Cleaning..."
	@go clean
	@rm -rf $(PROTO_GEN_DIR)/*/infrastructure/proto
	@rm -f ${BINARY_NAME}

run: proto wire
	@echo "Running app..."
	@go run ${MAIN_PACKAGE}

proto:
	@echo "Generating grpc code from proto files"
	$(foreach bc,$(BCS), \
		\
		mkdir -p $(CONTEXTS_DIR)/$(bc)/infrastructure/proto; \
		\
		$(PROTOC) --proto_path=$(PROTO_DIR) \
			--go_out=$(CONTEXTS_DIR)/$(bc)/infrastructure/proto \
			--go_opt=paths=source_relative \
			--go-grpc_out=$(CONTEXTS_DIR)/$(bc)/infrastructure/proto \
			--go-grpc_opt=paths=source_relative \
			$(PROTO_DIR)/$(bc).proto; \
		\
		$(shell go env GOPATH)/bin/$(PROTOC_GO_INJECT_TAG) \
			-input=./internal/contexts/$(bc)/infrastructure/proto/*.pb.go; \
		\
	)

.PHONY: all wire build clean run proto
