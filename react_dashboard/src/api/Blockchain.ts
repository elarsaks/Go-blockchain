import axios from "axios";

function fetchBlockchainData(): Promise<Blockchain> {
  return axios
    .get<Blockchain>("http://localhost:5001/")
    .then((response) => response.data);
}

export { fetchBlockchainData };
