build:
	@go build -o bin/out cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/out