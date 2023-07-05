# Use an official Go alpine image as the base
FROM golang:alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifest and download the dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code from the host machine to the container
COPY . .

# Build the Go application
RUN go build -o main .

# Set the entry point for the container
ENTRYPOINT ["./main"]
