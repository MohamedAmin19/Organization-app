# Use the official Go image as a base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the entire project directory into the container's workspace
COPY . .

# Download dependencies
RUN go mod download

# Build the Go app
RUN go build -o main cmd/main.go

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
