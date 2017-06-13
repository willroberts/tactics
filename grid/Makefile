check:
	errcheck ./...
	go vet ./...
	golint ./...

test:
	go test -v -coverprofile

test_graphics:
	go test -v -tags=graphics -coverprofile

coverage:
	go test -v -coverprofile cover.out
	go tool cover -func=cover.out

docs:
	@mkdir -p doc
	godoc github.com/willroberts/tactics/grid > doc/godoc.txt
