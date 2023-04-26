SHELL:=/bin/bash -O extglob
BINARY=api
VERSION=0.1.0
LDFLAGS=-ldflags "-X main.Version=${VERSION}"

build:
	go build ${LDFLAGS} -o ${BINARY} src/api/main.go

run:
	@go run src/api/main.go

up:
	docker-compose up --build

down:
	docker-compose down --remove-orphans