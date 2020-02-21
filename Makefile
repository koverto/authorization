.DEFAULT_GOAL := run

.PHONY: build
build: gen
	go build ./cmd/authorization

.PHONY: docker
docker: build
	docker build . -t koverto/authorization:latest

.PHONY: gen
gen:
	go generate ./api

.PHONY: run
run: gen
	go run ./cmd/authorization
