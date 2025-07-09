FROM golang:1.24.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o server .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .

RUN chmod +x server

CMD ["./server"]