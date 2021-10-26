MAIN=gorth
BIN=./bin/$(MAIN)
SAMPLE=./sample/test.gorth

.PHONY: ALL
all: deps test build

.PHONY: test
test:
	go test ./...

.PHONY: build
build:
	go build -o $(BIN) cmd/$(MAIN)/main.go

.PHONY: deps
deps:
	go mod tidy

init:
	go mod init

run: build
	$(BIN) -d -f $(SAMPLE)
