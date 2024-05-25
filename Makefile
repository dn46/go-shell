build:
	@go build -o bin/go-shell

run: build
	@./bin/go-shell

test:
	@go test -v ./...