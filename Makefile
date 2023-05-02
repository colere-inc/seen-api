.DEFAULT_GOAL := run

IMAGE_NAME=seen-api:latest
PORT=8080

tidy:
	go mod tidy

fmt: tidy
	goimports -local https://github.com/colere-inc/seen-api -w .

lint: fmt
	staticcheck

vet: lint
	go vet

build: vet
	docker build -t ${IMAGE_NAME} .
	docker image prune -f

run: build
	docker run --env-file .env -p ${PORT}:${PORT} --rm ${IMAGE_NAME}
