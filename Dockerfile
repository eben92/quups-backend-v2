# syntax=docker/dockerfile:1

FROM golang:1.22 as builder

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code including Go files, Makefile, and environment variables
COPY . ./

# Copy migrations

# Build the application
RUN CGO_ENABLED=0 GOOS=linux make build

# Start a new stage to reduce the final image size
FROM alpine:latest

# Copy the binary from the builder stage
COPY --from=builder /app/main /app/main

# Expose the port
EXPOSE 8080

# Set the working directory and command to run the application
WORKDIR /app
CMD ["./main"]
