FROM golang:1.24.1-alpine AS builder

LABEL author="Israel Schroeder Moreira"

LABEL maintainer="israelschroederm@gmail.com"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build cmd/server/main.go

FROM alpine:latest

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

RUN apk add --no-cache procps

WORKDIR /app

COPY --from=builder /app/main .

RUN chown -R appuser:appgroup /app

USER appuser

EXPOSE 3000

CMD ["./main"]
