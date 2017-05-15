check:
	@go vet ./...
	@golint ./...

test:
	@go test ./...

build:
	@mkdir -p bin
	@go build -o bin/tactics
