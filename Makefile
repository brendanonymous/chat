generate:
	@echo ""; echo "Generating fakes..."
	@go generate ./...

build:
	@go build -o bin/chat ./cmd

run: build
	@go run ./cmd/main.go

test:
	@echo ""; echo "Unit test report..."
	@go test ./... -coverprofile=c.out
	@echo ""; echo "Coverage report..."
	@go tool cover -func=c.out 