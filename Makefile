MAIN=gorth
BIN=./bin/$(MAIN)

.PHONY: ALL
all: test build

.PHONY: test
test: deps
	go test ./...

.PHONY: build
build: deps
	go build -o $(BIN) cmd/$(MAIN)/main.go

.PHONY: deps
deps:
	go mod tidy

init:
	go mod init

run: build
	$(BIN)