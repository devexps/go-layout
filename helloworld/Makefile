# Not Support Windows

GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)

API_PROTO_PATH=../api/proto
API_PROTO_FILES=$(shell cd $(API_PROTO_PATH) && find . -name *.proto)

# generate api proto
proto:
	@cd $(API_PROTO_PATH) && \
	mkdir -p ../gen ../gen/go ../gen/docs && \
	protoc --proto_path=./ \
		--proto_path=../third_party \
		--go_out=paths=source_relative:../gen/go \
		--go-http_out=paths=source_relative:../gen/go \
		--go-grpc_out=paths=source_relative:../gen/go \
		--openapi_out=fq_schema_naming=true,default_response=false:../gen/docs \
		$(API_PROTO_FILES)

# generate
generate:
	@go mod tidy
	@go generate ./...

# generate all
all:
	@make proto;
	@make generate;

# run tests
test:
	@go test -race ./...

# run coverage tests
cover:
	@go test -v ./... -coverprofile=coverage.out

# build golang application
build:
	@go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

# clean build files
clean:
	@go clean
	@rm -f "coverage.out"

# build service app
app: proto generate build

# run debug application
run:
	@go run ./cmd/server -conf ./configs

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help