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

# Build the Go app
RUN go build -o gss-backend ./cmd

# Expose port 3001
EXPOSE 3001

# Run the application
CMD ["./gss-backend"]