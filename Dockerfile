# Start from the official Golang image
FROM golang:1.23 AS builder

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Create app directory inside container
WORKDIR /app

# Copy go.mod and go.sum first, then download deps
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the Go app
RUN go build -o main .

# Command to run the executable
CMD ["./main"]
