# Use official Go image as base
FROM golang:1.24 as builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go app with CGO disabled
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .

# Use a lightweight image for the final container
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy built binary from builder
COPY --from=builder /app/app .

# Expose your app port (change if needed)
EXPOSE 8080

# Run the Go app
CMD ["./app"]
