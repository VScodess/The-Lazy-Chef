# Use the official Go image with the appropriate version as the base image
FROM golang:1.22-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the workspace
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the entire backend code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the application on port 8080 (or the port your application is listening on)
EXPOSE 8080

# Command to run the application
CMD ["./main"]
