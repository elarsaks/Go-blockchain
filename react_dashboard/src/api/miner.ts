import axios from "axios";

let { REACT_APP_GATEWAY_API_URL } = process.env;
if (!REACT_APP_GATEWAY_API_URL) {
  REACT_APP_GATEWAY_API_URL = "http://localhost:5000";
}

// Fetch latest blocks
function fetchBlockchainData(): Promise<[Block]> {
  return axios
    .get<[Block]>(REACT_APP_GATEWAY_API_URL + "/miner/blocks?amount=10")
    .then((response) => response.data);
}

// Fetch miner wallet details
function fetchMinerWalletDetails(minerId: string): Promise<WalletDetails> {
  return axios
    .post<WalletDetailsResponse>(
      REACT_APP_GATEWAY_API_URL + "/miner/wallet?miner_id=" + minerId
    )
    .then(({ data }) => {
      const camelCaseResponseData: WalletDetails = {
        blockchainAddress: data.blockchain_address,
        privateKey: data.private_key,
        publicKey: data.public_key,
      };

      return camelCaseResponseData;
    });
}

export { fetchBlockchainData, fetchMinerWalletDetails };
