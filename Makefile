BINARY_NAME=library

MAIN_PACKAGE=./cmd/api

WIRE=./vendor/github.com/google/wire/cmd/wire

BCS := authentication borrow
PROTO_DIR := ./proto
CONTEXTS_DIR := ./internal/contexts
CONTEXT := borrow

PROTOC := protoc
PROTOC_GEN_GO := protoc-gen-go
PROTOC_GEN_GO_GRPC := protoc-gen-go-grpc
PROTOC_GO_INJECT_TAG := protoc-go-inject-tag

.PHONY: all wire build clean run proto migration

wire:
	@echo "Generating dependency injection..."
	@go generate ./...

build: proto wire
	@echo "Building app..."
	@go build -o ${BINARY_NAME} ${MAIN_PACKAGE}

clean:
	@echo "Cleaning..."
	@go clean
	@go clean -testcache
	@rm -rf $(PROTO_GEN_DIR)/*/infrastructure/proto
	@rm -f ${BINARY_NAME}

run: proto wire
	@echo "Running app..."
	@go run ${MAIN_PACKAGE}

test: proto wire
	@echo "Running tests..."
	@go test ./... | grep -v "no test files"

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

migration:
	@echo "Generating migration..."
	@if [ -z $(MG_NAME) ]; then \
		echo "Please provide a migration name (MG_NAME)"; \
		exit 1; \
	fi

	@if [ -z $(CONTEXT) ]; then \
		echo "Undefined CONTEXT, defaulting to \"borrow\""; \
		exit 1; \
	fi

	@migrate  create -ext sql -dir ./internal/contexts/$(CONTEXT)/infrastructure/sql/migrations/ -seq $(MG_NAME)

