import axios from "axios";
import { Blockchain } from "../Type";

async function fetchBlockchainData(): Promise<Blockchain> {
  try {
    const response = await axios.get<Blockchain>("http://localhost:5001/");
    console.log("Fetched blockchain data:", response);
    return response.data;
  } catch (error) {
    console.error("Failed to fetch blockchain data:", error);
    return [];
  }
}

export { fetchBlockchainData };
