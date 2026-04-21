# ---- Build stage ----
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Install git (needed for go mod)
RUN apk add --no-cache git

# Cache dependencies first
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .

# ---- Runtime stage ----
FROM alpine:3.19

WORKDIR /app

# Add CA certs (important for real-world apps)
RUN apk add --no-cache ca-certificates

# Copy binary from builder
COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]
