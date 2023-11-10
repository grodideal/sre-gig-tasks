# Use the official Go image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download and install dependencies
RUN go mod download

# Copy the rest of the application code to the container
COPY . .

# Build the Go application inside the container
RUN go build -o main .

# Expose port 8080 for the web server
EXPOSE 8080

# Command to run the application when the container starts
CMD ["./main"]