# Step 1: Use official Golang image for building
FROM golang:1.23.5 AS builder

WORKDIR /app

RUN apt-get update && apt-get install -y curl && \
    curl -fsSL https://deb.nodesource.com/setup_18.x | bash - && \
    apt-get install -y nodejs && \
    npm install -g npm

COPY go.mod go.sum ./

RUN go mod tidy && go mod download

RUN go install github.com/a-h/templ/cmd/templ@latest

COPY . .

RUN npm install -D tailwindcss

RUN npx tailwindcss -i view/css/app.css -o ./view/public/styles.css

RUN templ generate view

RUN go build -tags prod -o bin/main ./cmd/app

FROM gcr.io/distroless/base-debian11

WORKDIR /app

COPY --from=builder /app/bin/main /app/main

EXPOSE 8080

CMD ["/app/main"]
