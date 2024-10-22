build:
	@go build -o bin/goapi cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/goapi