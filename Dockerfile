# Stage 1: Build Go web server
FROM golang:1.16 AS builder

WORKDIR /app

# Copy Go source code
COPY . .

# Install Go dependencies
RUN go mod init example.com/m/v2
RUN go mod tidy

# Build the Go binary
RUN go build -o bin/host main.go

# Expose the port for the web server
EXPOSE ${GO_MILVUS_PORT}

# Start the web server
CMD ["./bin/host"]
