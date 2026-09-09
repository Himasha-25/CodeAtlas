.PHONY: run build test lint

run:
	go run cmd/server/main.go

build:
	go build -o bin/server ./cmd/server

test:
	go test ./...

test-analyzer:
	go test ./analyzer/...

lint:
	golangci-lint run
