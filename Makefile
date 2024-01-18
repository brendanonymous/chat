generate:
	@go generate ./...

build:
	@go build -o bin/chat ./cmd

run: build
	go run ./cmd/main.go

test:
	go test -v ./...