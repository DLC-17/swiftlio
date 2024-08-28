run: build
	@./bin/app

test:
	@go test -v ./...

build: 
	@go build -o bin/app .

tailwind:
	@tailwindcss -i view/css/styles.css -o public/styles.css --watch

templ:
	@templ generate --watch --proxy=http://localhost:6969
