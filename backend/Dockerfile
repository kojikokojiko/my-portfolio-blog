# Build stage
FROM golang:1.22-alpine AS builder

# Install necessary packages
RUN apk update && apk add --no-cache git

WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY backend/go.mod backend/go.sum ./
RUN go mod download

# Copy the application code
COPY backend/ ./

# Build the application
RUN go build -o main .

# Development stage
FROM golang:1.22-alpine AS development

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Install air for hot reloading
RUN go install github.com/cosmtrek/air@latest

# Copy .air.toml configuration
COPY backend/.air.toml /app/.air.toml

# Copy the application code for hot reloading
COPY backend/ ./

# Command to run air for hot reloading
CMD ["air", "-c", ".air.toml"]

# Production stage
FROM alpine:latest AS production

WORKDIR /app

# Copy the built binary and .env from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Command to run the built application
CMD ["./main"]
