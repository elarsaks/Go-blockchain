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

function fetchWalletData(): Promise<Wallet> {
  return axios
    .get<Wallet>("http://localhost:5001/wallet")
    .then((response) => response.data)
    .catch((error) => {
      console.error("Failed to fetch wallet data:", error);
      return { blockchainAddress: "", privateKey: "", publicKey: "" };
    });
}



export { fetchBlockchainData, fetchWalletData };
