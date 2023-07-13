import axios from "axios";

function fetchBlockchainData(): Promise<[Block]> {
  return axios
    .get<[Block]>("http://localhost:5001/last10") // TODO: this should be docker container name
    .then((response) => response.data);
}

export { fetchBlockchainData };
