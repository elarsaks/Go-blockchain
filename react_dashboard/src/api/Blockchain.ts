import axios from "axios";
import { Blockchain } from "../Type";

function fetchBlockchainData(): Promise<Blockchain> {
  return axios
    .get<Blockchain>("http://localhost:5001/")
    .then((response) => response.data)
    .catch((error) => {
      console.error("Failed to fetch blockchain data:", error);
      return { chain: [] };
    });
}


export { fetchBlockchainData };
