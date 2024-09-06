FROM golang:1.23-alpine AS go-builder

RUN apk add --no-cache make curl

WORKDIR /app

RUN curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-arm64 \
    && chmod +x tailwindcss-linux-arm64 \
    && mv tailwindcss-linux-arm64 tailwindcss \
    && mv tailwindcss /usr/local/bin/

COPY go.mod go.sum ./

RUN go mod download && go mod tidy

RUN go install github.com/air-verse/air@latest \
    && go install github.com/a-h/templ/cmd/templ@latest

COPY . .
