import axios from "axios";

function fetchUserWalletDetails(): Promise<WalletDetails> {
  return axios
    .post<WalletDetailsResponse>("http://localhost:5000/wallet") // TODO: Rename api endpoint
    .then(({ data }) => {
      const camelCaseResponseData: WalletDetails = {
        blockchainAddress: data.blockchain_address,
        privateKey: data.private_key,
        publicKey: data.public_key,
      };

      return camelCaseResponseData;
    });
}

function fetchMinerWalletDetails(): Promise<WalletDetails> {
  return axios
    .post<WalletDetailsResponse>("http://localhost:5000/miner-wallet-details") // TODO: API endpoint
    .then(({ data }) => {
      const camelCaseResponseData: WalletDetails = {
        blockchainAddress: data.blockchain_address,
        privateKey: data.private_key,
        publicKey: data.public_key,
      };

      return camelCaseResponseData;
    });
}

export { fetchMinerWalletDetails, fetchUserWalletDetails };
