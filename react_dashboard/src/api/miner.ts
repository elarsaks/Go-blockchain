import axios from "axios";

const { REACT_APP_GATEWAY_API_URL } = process.env;
const WALLET_SERVER_URL = REACT_APP_GATEWAY_API_URL
  ? REACT_APP_GATEWAY_API_URL
  : "https://go-blockchain.azurewebsites.net"; // During build there is no env variables

// Fetch latest blocks
function fetchBlockchainData(): Promise<[Block]> {
  return axios
    .get<[Block]>(WALLET_SERVER_URL + "/miner/blocks?amount=10")
    .then((response) => response.data);
}

// Fetch miner wallet details
function fetchMinerWalletDetails(minerAdress: string): Promise<WalletDetails> {
  return axios
    .post<WalletDetailsResponse>(minerAdress + "/miner/wallet")
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
