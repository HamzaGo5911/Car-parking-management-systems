# Base image
FROM golang:1.17-alpine as builder

# Set working directory
WORKDIR /app

# Copy go modules
COPY go.mod go.sum ./

# Download go modules
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN go build -o car-parking-management-systems

# Final image
FROM alpine:3.14

# Install dependencies
RUN apk add --no-cache ca-certificates

# Copy binary from builder stage
COPY --from=builder /app/car-parking-management-systems /usr/local/bin/car-parking-management-systems
# Sets the executable permission
RUN chmod +x /usr/local/bin/car-parking-management-systems

# Set command
CMD ["/usr/local/bin/car-parking-management-systems"]
