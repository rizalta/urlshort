templ:
	@templ generate

tw:
	@npx tailwindcss -i ./web/assets/css/input.css -o ./web/assets/css/output.css --minify

build: templ tw
	@go build -o ./bin/urlshort ./cmd/urlshort

run: build
	./bin/urlshort

PHONY: templ tw build run
