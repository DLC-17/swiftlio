# --- Loading .env vars ---
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# --- Make targets ---
.PHONY: dev

air:
	@air

build: 
	@go build -o bin/swiftlio ./cmd/main.go

tailwind:
	@tailwindcss -i ./ui/css/styles.css -o ./ui/public/styles.css --watch

templ:
	@templ generate --watch --proxy="http://localhost:6969" --proxybind="0.0.0.0" --proxyport=7331 --open-browser=false

test:
	@go test -v ./...

dev:
	make -j3 air templ tailwind 

db-up:
	@cd internal/sql/schema && goose postgres $(DATABASE_URL) up

db-down:
	@cd internal/sql/schema && goose postgres $(DATABASE_URL) down
