provider "docker" {
  host = "unix:///var/run/docker.sock" # Use the Docker socket on your system
}

resource "docker_image" "react_dashboard_image" {
  name = "react-dashboard-image" # Replace with your actual image name
  build {
    context = "./cmd/react_dashboard"
  }
}

resource "docker_container" "react_dashboard" {
  name  = "react-dashboard"
  image = docker_image.react_dashboard_image.name
  ports {
    internal = 3000
    external = 3000
  }
  environment = {
    REACT_APP_GATEWAY_API_URL = "http://localhost:${var.WALLET_SERVER_PORT}"
  }
  volumes {
    host_path      = "./cmd/react_dashboard"
    container_path = "/app"
  }
  volumes {
    host_path      = "./cmd/react_dashboard/node_modules"
    container_path = "/app/node_modules"
  }
  restart = "unless-stopped"
}

resource "docker_container" "wallet_server" {
  name  = "wallet-server"
  image = "your-wallet-server-image" # Replace with your actual image name
  ports {
    internal = 5000
    external = 5000
  }
  environment = {
    PORT       = var.WALLET_SERVER_PORT
    MINER_HOST = "miner-1"
  }
  volumes {
    host_path      = "."
    container_path = "/app"
  }
  restart = "unless-stopped"
}

# Define a list of miner names
variable "miner_names" {
  default = ["miner-1", "miner-2", "miner-3"]
}

# Create miners using a for loop
resource "docker_container" "miner" {
  count = length(var.miner_names)

  name  = var.miner_names[count.index]
  image = "your-blockchain-server-image" # Replace with your actual image name
  ports {
    internal = 5000 + count.index
    external = 5000 + count.index
  }
  environment = {
    COMMAND    = "air"
    PORT       = tostring(5000 + count.index)
    MINER_HOST = "miner"
  }
  volumes {
    host_path      = "."
    container_path = "/app"
  }
  restart = "unless-stopped"
}
