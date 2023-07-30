# Start from the base Golang image
FROM golang:1.18.0-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Install the "air" tool for live-reloading (if needed)
RUN go install github.com/cosmtrek/air@latest

# Copy the entire project into the container
COPY . .

# Set the working directory to "wallet_server"
WORKDIR /app/cmd/wallet_server

# Set the port number to expose for HTTP (port 8080)
EXPOSE ${PORT}

# Set the default command to run the application
CMD ["sh", "-c", "${COMMAND}"]
