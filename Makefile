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

execute:
	EXEC_IN=/home/fcordier/iexecdev/predigo IEXEC_DATASET_FILENAME=ds.csv IEXEC_OUT=/home/fcordier/iexecdev/predigo ./salary-estimator DEV_NODEJS PARIS BACHELOR 14

