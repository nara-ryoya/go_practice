:DEFAULT_GOAL := build

fmt:
	go fmt
.PHONY: fmt

lint: fmt
		staticcheck
.PHONY: liint
vet: fmt
		go vet ./...
			shadow ./...
.PHONY: vet

build: vet
		go mod tidy
		go build
.PHONY: build