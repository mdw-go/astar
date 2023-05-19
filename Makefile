#!/usr/bin/make -f

test: fmt build
	go test -v -short ./...

build:
	go build ./...

fmt:
	go fmt ./... && go mod tidy
