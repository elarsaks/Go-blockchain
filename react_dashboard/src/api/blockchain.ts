import axios from "axios";

function fetchBlockchainData(): Promise<[Block]> {
  return axios
    .get<[Block]>("http://localhost:5000/miner/blocks?amount=10")
    .then((response) => response.data);
}

export { fetchBlockchainData };
