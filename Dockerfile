# Start from a Debian-based image with Go installed
FROM golang:1.18 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod .
# COPY go.sum .

# Download and install dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o app main.go

# Start a new stage from a lightweight base image
FROM debian:buster-slim

# Set the working directory inside the container
WORKDIR /app

# Copy the built executable from the previous stage
COPY --from=build /app/app .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./app"]