import axios from "axios";
import { Blockchain, Wallet } from "../Type";

//? Is there need for a util folder?
function snakeToCamelCase(snakeCaseString: string): string {
  return snakeCaseString.replace(/(_\w)/g, (match) => match[1].toUpperCase());
}

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
