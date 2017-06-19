.PHONY: docs

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

html_coverage:
	@go test -v -coverprofile cover.out
	@go tool cover -html=cover.out
	@rm cover.out

docs:
	@mkdir -p docs
	@godoc github.com/willroberts/tactics/engine > docs/engine.txt
	@godoc github.com/willroberts/tactics/engine/input > docs/input.txt
	@godoc github.com/willroberts/tactics/engine/menu > docs/menu.txt
	@godoc github.com/willroberts/tactics/game/scenes > docs/scenes.txt
	@godoc github.com/willroberts/tactics/game/unit > docs/unit.txt
	@godoc github.com/willroberts/tactics/grid > docs/grid.txt
	@godoc github.com/willroberts/tactics/tmx > docs/tmx.txt

build:
	@mkdir -p bin
	go build -o bin/tactics
