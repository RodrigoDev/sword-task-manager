.PHONY: build run test

PROJECT?=sword-task-manager

default: build

build: test build-local

build-local:
	go build -o ./app ./cmd/server

run: build
	./app

run-local: build-local
	./app

build-docker: test
	docker-compose build

run-docker: build-docker
	docker-compose up -d

stop-docker:
	docker-compose stop

clean-docker:
	docker-compose rm -s -f

lint: get-linter
	golangci-lint run --timeout=5m

get-linter:
	command -v golangci-lint || curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b ${GOPATH}/bin

test: lint
	go fmt ./...
	go test -vet all ./...

test-race:
	go test -v -race ./...

test-integration:
	go test -v -vet all -tags=integration ./... -coverprofile=integration.out

test-all: test test-race test-integration

cover-ci:
	go tool cover -func=integration.out

cover:
	echo unit tests only
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

mod:
	go mod vendor -v

tidy:
	go mod tidy

docs:
	godoc -http=:6060
