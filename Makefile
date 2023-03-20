.PHONY: build

default:build

build:
	CGO_ENABLED=0 go build  --ldflags '-s' -buildmode=pie

start:
	go run main.go

format:
	go fmt ./...

lint:
	golangci-lint run

tidy:
	go mod tidy

