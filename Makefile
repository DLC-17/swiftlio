.PHONY: dev

air:
	@air

build: 
	@go build -o bin/swiftlio ./cmd/main.go

tailwind:
	@tailwindcss -i ./ui/css/styles.css -o ./ui/public/styles.css --watch

templ:
	@templ generate --watch --proxy="http://localhost:6969" --open-browser=false

test:
	@go test -v ./...

dev:
	make -j3 air templ tailwind
