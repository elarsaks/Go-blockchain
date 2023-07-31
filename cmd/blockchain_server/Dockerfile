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

# Copy the entire parent_folder into the container
COPY . .

# Set the working directory to "blockchain_server"
WORKDIR /app/cmd/blockchain_server

# Set the port number to expose
EXPOSE ${BLOCKCHAIN_SERVER_PORT}

# Set the default command to run the application
CMD ["sh", "-c", "air run"]
