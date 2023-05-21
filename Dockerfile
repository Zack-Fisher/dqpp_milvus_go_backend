# Stage 1: Build Go web server
FROM golang:1.16 AS builder

WORKDIR /app

# Copy Go modules manifests
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy Go source code
COPY . .

# Build the Go binary
RUN go build -o webserver .

# Copy the Go binary from the builder stage
COPY --from=builder /app/webserver /app/webserver

# Expose the port for the web server
EXPOSE 8080

# Start the web server
CMD ["/app/webserver"]
