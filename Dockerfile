# Build stage
FROM golang:1.24.5 AS builder

WORKDIR /app

# Copy go mod files first (for caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o app

# Run stage (minimal image)
FROM alpine:latest

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/app .

# Expose port (if your app uses one)
EXPOSE 8080

# Run binary
CMD ["./app"]
