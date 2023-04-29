.DEFAULT_GOAL := build

tidy:
	go mod tidy

fmt: tidy
	goimports -local https://github.com/colere-inc/seen-api -w .

lint: fmt
	staticcheck

vet: lint
	go vet

build: vet
	go build
