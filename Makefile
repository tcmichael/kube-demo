all: build
.PHONY: all

build:
	go mod vendor
	GO111MODULE="on" GOOS=linux go build -o bin/webhook pkg/cmd/webhook/main.go
.PHONY: build