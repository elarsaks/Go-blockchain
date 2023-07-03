import axios from "axios";
import { Blockchain, Wallet } from "../Type";

//? Is there need for a util folder?
function snakeToCamelCase(snakeCaseString: string): string {
  return snakeCaseString.replace(/(_\w)/g, (match) => match[1].toUpperCase());
}

function fetchWalletData(): Promise<Wallet> {
  return axios
    .post<Wallet>("http://localhost:5000/wallet")
    .then((response) => response.data)
    .catch((error) => {
      console.error("Failed to fetch wallet data:", error);
      return { blockchainAddress: "", privateKey: "", publicKey: "" };
    });
}

export { fetchWalletData };
