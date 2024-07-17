.DEFAULT_GOAL := build
.SILENT:

.PHONY: build
build:
	go build -o ./build/mkobj ./cmd/main.go
