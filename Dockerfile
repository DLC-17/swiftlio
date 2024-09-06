# Stage 1: Golang Build Stage
FROM golang:1.23-alpine AS go-builder

# Install make utility
RUN apk add --no-cache make
RUN apk add --no-cache curl 

WORKDIR /app

# Installing tailwindcss
RUN curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-arm64 \
    && chmod +x tailwindcss-linux-arm64 \
    && mv tailwindcss-linux-arm64 tailwindcss \
    && mv tailwindcss /usr/local/bin/

# Copy go.mod and go.sum first to leverage Docker cache for dependency download
COPY go.mod go.sum ./

# Install Go dependencies
RUN go mod download \ 
    && go mod tidy

# Install additional Go tools (air and templ)
RUN go install github.com/air-verse/air@latest \
    && go install github.com/a-h/templ/cmd/templ@latest

COPY . .

# Expose the necessary port
EXPOSE 6969
EXPOSE 7331

CMD ["make", "dev"]
