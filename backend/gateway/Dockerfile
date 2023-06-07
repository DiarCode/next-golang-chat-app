# Dockerfile for the gateway

# Use the official Golang image as the base image
FROM golang:1.19-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o gateway_build

# Run the gateway when the container starts
CMD ["/app/gateway_build"]