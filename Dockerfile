FROM golang:latest

WORKDIR /app

COPY . /app

RUN go mod download

CMD go run ./cmd/heimdallr/*.go