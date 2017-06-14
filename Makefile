check:
	errcheck ./...
	go vet ./...
	golint ./...

test:
	go test -v ./...

coverage:
	@go test -v -coverprofile cover.out
	@go tool cover -func=cover.out
	@rm cover.out

docs:
	@godoc github.com/willroberts/tactics/engine > engine/doc.txt
	@godoc github.com/willroberts/tactics/grid > grid/doc.txt
	@godoc github.com/willroberts/tactics/tmx > tmx/doc.txt
	@godoc github.com/willroberts/tactics/unit > unit/doc.txt

build:
	@mkdir -p bin
	go build -o bin/tactics
