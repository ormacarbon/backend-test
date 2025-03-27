FROM golang:1.24.1-alpine3.21

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY . .
RUN go mod download

CMD ["air", "run"]