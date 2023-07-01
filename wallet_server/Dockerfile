FROM golang:1.18.0-alpine

WORKDIR /app/wallet_server

RUN go install github.com/cosmtrek/air@latest

CMD ["sh", "-c", "PORT=5001 air"]
