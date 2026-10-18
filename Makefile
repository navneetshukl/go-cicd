.PHONY: fmt vet test cover build run

fmt:
	gofmt -l -w .

vet:
	go vet ./...

test:
	go test -race -count=1 ./...

cover:
	go test -race -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out | tail -n 1

build:
	CGO_ENABLED=0 go build -o bin/api ./cmd/api

run:
	go run ./cmd/api