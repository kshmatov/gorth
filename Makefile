MAIN=gorth
BIN=./bin/$(MAIN)
SAMPLE=./sample/test.gorth

.PHONY: ALL
all: deps test build

.PHONY: vet
vet:
	go vet ./...

.PHONY: test
test: vet
	go test ./...

.PHONY: build
build: vet
	go build -o $(BIN) cmd/$(MAIN)/main.go

.PHONY: deps
deps:
	go mod tidy

init:
	go mod init

run: build
	$(BIN) -d -f $(SAMPLE)
