# Tell Makefile to use bash
SHELL := /bin/bash

.PHONY: init up up-dev down install install-dev cache clear key seed bash

build:
	go env -w GOPRIVATE=github.com
	go mod tidy

swag:
	swag init

test:
	go test -v ./service/...

start:
	gin -p 8000 -a 8001 main.go