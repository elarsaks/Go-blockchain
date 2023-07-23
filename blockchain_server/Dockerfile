FROM golang:1.18.0-alpine

WORKDIR /app/blockchain_server

RUN go install github.com/cosmtrek/air@latest

EXPOSE ${BLOCKCHAIN_SERVER_PORT}

CMD ["sh", "-c", "${COMMAND}"]
