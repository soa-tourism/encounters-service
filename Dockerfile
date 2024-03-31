# Start from a Golang base image
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /go/src/encounters-service

# Initialize Go module
RUN go mod init encounters-service

# Copy the source code from the current directory to the Working Directory inside the container
COPY src .

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8090

# Command to run the executable
CMD ["./main"]
