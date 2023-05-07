.PHONY: default build clean tests

BUILD_DIR := build

default: build

build:
	go build -ldflags="-s -w" -o $(BUILD_DIR) ./...
	strip -S $(BUILD_DIR)/*
clean:
	rm -f $(BUILD_DIR)/*
tests:
	go test -cover ./pkg/...

