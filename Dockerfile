FROM golang:1.23.6 as builder

# Define build env
ENV GOOS=linux
ENV CGO_ENABLED=0

# Add a work directory
WORKDIR /app

# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy app files
COPY . .

# Build app
RUN go build -o app cmd/main.go

FROM alpine:3.14 as production

# Add certificates
RUN apk add --no-cache ca-certificates

# Copy built binary from builder
COPY --from=builder /app/app /usr/local/bin/app

# Copy the .env file
COPY --from=builder /app/.env /.env
# Or if you prefer to keep it in the app directory:
# COPY --from=builder /app/.env /app/.env

# Add a non-root user and set ownership
RUN adduser -D appuser && \
    chown appuser /usr/local/bin/app && \
    chmod +x /usr/local/bin/app && \
    chown appuser /.env
# If you used the alternative path above, use this instead:
# chown appuser /app/.env

# Switch to non-root user
USER appuser

# Expose port
EXPOSE 8000

# Exec built binary
CMD ["/usr/local/bin/app"]