generate:
	@go generate ./...

build:
	@go build -o bin/boat-auction ./cmd

run: build
	go run ./cmd/main.go

test:
	go test -v ./...