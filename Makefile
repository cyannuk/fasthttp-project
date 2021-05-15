build:
	go build -v -ldflags="-w -s"

dependencies:
	go mod download

.PHONY: build dependencies
.DEFAULT_GOAL := build
