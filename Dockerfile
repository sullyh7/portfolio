# Step 1: Use latest Golang 1.23.5 for building
FROM golang:1.23.5 AS builder

# Set working directory
WORKDIR /app

# Install Node.js & npm (needed for Tailwind CSS)
RUN apt-get update && apt-get install -y curl && \
    curl -fsSL https://deb.nodesource.com/setup_18.x | bash - && \
    apt-get install -y nodejs && \
    npm install -g npm

# Copy Go module files first to leverage Docker's cache
COPY go.mod go.sum ./

# Download Go dependencies
RUN go mod tidy && go mod download

# Install Templ
RUN go install github.com/a-h/templ/cmd/templ@latest

# Copy the rest of the project files
COPY . .

# Install Tailwind CSS
RUN npm install -D tailwindcss

# Build CSS using Tailwind
RUN npx tailwindcss -i view/css/app.css -o ./view/public/styles.css

# Generate Templ templates
RUN templ generate view

# Build the Go application
RUN go build -tags prod -o bin/main ./cmd/app

# Step 2: Use latest Debian as runtime (Fixes GLIBC issues)
FROM debian:latest

# Set working directory
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/bin/main /app/main

# Expose the port (update if needed)
EXPOSE 8080

# Run the application
CMD ["/app/main"]
