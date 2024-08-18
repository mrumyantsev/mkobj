.DEFAULT_GOAL := build
.SILENT:

.PHONY: build
build:
	go build -o ./build/mkobj ./cmd/mkobj/...

.PHONY: install
install:
	cp ./build/mkobj /usr/local/bin
