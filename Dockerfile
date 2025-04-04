FROM golang:1.24.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o backend-test ./cmd/main.go

FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=builder /app/backend-test .

EXPOSE 8080

CMD ["./backend-test"]
