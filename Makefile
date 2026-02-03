BINARY_NAME=certinfo
BIN_DIR=bin

.PHONY: all build build-linux build-freebsd clean help

help:
	@echo Usage: make [target]
	@echo.
	@echo Targets:
	@echo   all             Build for all platforms
	@echo   build           Build for local OS
	@echo   build-linux     Build for Linux AMD64
	@echo   build-freebsd   Build for FreeBSD AMD64
	@echo   clean           Remove build artifacts

all: build-linux build-freebsd

build:
	go build -o $(BINARY_NAME) main.go

build-linux:
	@go env -w GOOS=linux GOARCH=amd64
	go build -o bin/certinfo-linux main.go
	@go env -u GOOS GOARCH

build-freebsd:
	@go env -w GOOS=freebsd GOARCH=amd64
	go build -o bin/certinfo-freebsd main.go
	@go env -u GOOS GOARCH

clean:
	rm -rf $(BIN_DIR)
	rm -f $(BINARY_NAME)