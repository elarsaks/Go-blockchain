import axios from "axios";

const { REACT_APP_GATEWAY_API_URL } = process.env;
const WALLET_SERVER_URL = REACT_APP_GATEWAY_API_URL
  ? REACT_APP_GATEWAY_API_URL
  : "https://go-blockchain.azurewebsites.net"; // During build there is no env variables

function fetchUserWalletDetails(): Promise<WalletDetails> {
  return axios
    .post<WalletDetailsResponse>(WALLET_SERVER_URL + "/wallet")
    .then(({ data }) => {
      // console.log('User Details', data);
      const camelCaseResponseData: WalletDetails = {
        blockchainAddress: data.blockchain_address,
        privateKey: data.private_key,
        publicKey: data.public_key,
      };

      return camelCaseResponseData;
    });
}

function fetchWalletBalance(blockchainAddress: string): Promise<string> {
  return axios
    .get<BalanceResponse>(
      `${WALLET_SERVER_URL}/wallet/balance?blockchain_address=${blockchainAddress}`
    )
    .then(({ data }) => {
      if (data.error) {
        throw new Error(data.error);
      }
      return data.balance;
    });
}

function transaction(transaction: Transaction): Promise<string> {
  console.log(transaction);
  // Why this string ends up in golang as a number is beyond me
  return axios
    .post<string>(`${WALLET_SERVER_URL}/transaction`, transaction)
    .then(({ data }) => data);
}

export { fetchUserWalletDetails, fetchWalletBalance, transaction };
