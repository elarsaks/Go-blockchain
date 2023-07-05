import axios from "axios";

function fetchBlockchainData(): Promise<Blockchain> {
  return axios
    .get<Blockchain>("http://localhost:5005/")
    .then((response) => response.data);
}

export { fetchBlockchainData };
