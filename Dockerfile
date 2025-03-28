# Setting up my base image
FROM golang:1.24.1-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Get all dependencies
RUN go mod download

# Copy source code
COPY . .

# Expose port 3000
EXPOSE 3000

# Run the application
CMD ["go", "run", "cmd/main.go"]