check:
	errcheck ./...
	go vet ./...
	golint ./...

test:
	go test -v -coverprofile cover.out

coverage:
	go test -v -coverprofile cover.out
	go tool cover -func=cover.out

docs:
	@mkdir -p doc
	godoc github.com/willroberts/tactics/grid > doc/godoc.txt

build:
	@mkdir -p bin
	go build -o bin/tactics
