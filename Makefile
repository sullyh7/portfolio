run: build
	@./bin/main

css:
	@npx @tailwindcss/cli -i view/css/app.css -o ./view/public/styles.css


templ:
	@templ generate --watch --proxy=http://localhost:3000

build: css
	@templ generate view
	@go build -tags prod -o bin/main ./cmd/app

build-air: css
	@templ generate view
	@go build -tags prod -o tmp/main ./cmd/app

install:
	@go install github.com/a-h/templ/cmd/templ@latest
	@templ generate
	@go mod tidy
	@go mod download
	@npm install -D tailwindcss

