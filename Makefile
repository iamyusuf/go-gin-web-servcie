run: build
	@./bin/myservice
build:
	@go build -o bin/myservice .
