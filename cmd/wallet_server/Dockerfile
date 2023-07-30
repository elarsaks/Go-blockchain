# Start from the golang base image
FROM golang:1.18.0-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download 

# Copy the entire project
COPY . .

# Install the "air" tool for live-reloading (if needed)
RUN go install github.com/cosmtrek/air@latest

# Set the working directory to "wallet_server"
WORKDIR /app/cmd/wallet_server

# Set the port number to expose for HTTP (port 8080)
EXPOSE ${PORT}

# Set the default command to run the application
CMD ["sh", "-c", "${COMMAND}"]
