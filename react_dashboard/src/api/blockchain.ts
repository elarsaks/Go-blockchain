import axios from "axios";

const { REACT_APP_GATEWAY_API_URL } = process.env;
const WALLET_SERVER_URL = REACT_APP_GATEWAY_API_URL
  ? REACT_APP_GATEWAY_API_URL
  : "goblockchain.azurecr.io"; // During build there is no env variables

function fetchBlockchainData(): Promise<[Block]> {
  return axios
    .get<[Block]>(WALLET_SERVER_URL + "/miner/blocks?amount=10")
    .then((response) => response.data);
}

export { fetchBlockchainData };
