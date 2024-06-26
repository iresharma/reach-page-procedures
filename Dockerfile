# Stage 1: Build the Go application
FROM golang:1.19-alpine AS builder

# Install git and other dependencies
RUN apk update && apk add --no-cache git

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Set the working directory for the build context
WORKDIR /app/cmd

# Build the Go application
RUN go build -o /grpc-app .

# Stage 2: Create a smaller image with just the built binary
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /grpc-app .

# Inject environment variable from railway
ARG POSTGRES

# Set environment variables
ENV POSTGRES=$POSTGRES

# Expose the port the application runs on (assuming default gRPC port 50051)
EXPOSE 4001

# Command to run the application
CMD ["./grpc-app"]
