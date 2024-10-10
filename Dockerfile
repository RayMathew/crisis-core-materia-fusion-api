# Use the official Go image to build the application
FROM golang:1.23.0 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download all the dependencies (requires go.mod and go.sum)
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app with correct OS and architecture, and ensure executable permissions
RUN GOOS=linux GOARCH=amd64 go build -o api ./api && chmod +x api

# Start a new stage from scratch
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/api .

# Copy the certificates folder to the container
COPY certs /app/certs

# Expose the port the app runs on
EXPOSE 4444

# Command to run the executable
CMD ["./api"]
