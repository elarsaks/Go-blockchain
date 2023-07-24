FROM golang:1.18.0-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the entire parent_folder into the container
COPY . .

# Install the "air" tool for live-reloading (if needed)
RUN go install github.com/cosmtrek/air@latest

# Set the working directory to "wallet_server"
WORKDIR /app/wallet_server

# Set the port number to expose for HTTP (port 8080)
EXPOSE 8080

# Set the port number to expose for HTTPS (port 443)
EXPOSE 443

# Set the default command to run the application
CMD ["sh", "-c", "${COMMAND}"]
