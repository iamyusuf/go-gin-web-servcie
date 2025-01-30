run: build
	@./bin/myservice
build:
	@go build -o bin/myservice ./cmd/server/main.go
