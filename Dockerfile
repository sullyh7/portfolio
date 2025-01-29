# Step 1: Use official Golang image
FROM golang:1.23.5 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy && go mod download

RUN go install github.com/a-h/templ/cmd/templ@latest && \
    npm install -D tailwindcss

COPY . .

RUN npx @tailwindcss/cli -i view/css/app.css -o ./view/public/styles.css

RUN templ generate view

RUN go build -tags prod -o bin/main ./cmd/app

FROM gcr.io/distroless/base-debian11

WORKDIR /app

COPY --from=builder /app/bin/main /app/main

EXPOSE 8080

CMD ["/app/main"]
