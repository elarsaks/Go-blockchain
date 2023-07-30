FROM golang:1.18.0-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the entire parent_folder into the container
COPY ../../ .

# Install the "air" tool for live-reloading (if needed)
RUN go install github.com/cosmtrek/air@latest

# Set the working directory to "blockchain_server"
WORKDIR /app/cmd/blockchain_server

# Set the port number to expose
EXPOSE ${BLOCKCHAIN_SERVER_PORT}

# Set the default command to run the application
CMD ["sh", "-c", "${COMMAND}"]

