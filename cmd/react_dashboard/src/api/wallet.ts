import axios from "axios";

let { REACT_APP_GATEWAY_API_URL } = process.env;
if (!REACT_APP_GATEWAY_API_URL) {
  REACT_APP_GATEWAY_API_URL = "http://localhost:5000";
}

function fetchUserWalletDetails(): Promise<WalletDetails> {
  return axios
    .post<WalletDetailsResponse>(REACT_APP_GATEWAY_API_URL + "/user/wallet")
    .then(({ data }) => {
      console.log(data);

      const camelCaseResponseData: WalletDetails = {
        blockchainAddress: data.blockchainAddress,
        privateKey: data.privateKey,
        publicKey: data.publicKey,
      };

      return camelCaseResponseData;
    });
}

function fetchWalletBalance(blockchainAddress: string): Promise<string> {
  return axios
    .get<BalanceResponse>(
      `${REACT_APP_GATEWAY_API_URL}/wallet/balance?blockchainAddress=${blockchainAddress}`
    )
    .then(({ data }) => {
      if (data.error) {
        throw new Error(data.error);
      }
      return data.balance;
    });
}

function transaction(transaction: Transaction): Promise<any> {
  // Why this string ends up in golang as a number is beyond me
  return axios
    .post<any>(`${REACT_APP_GATEWAY_API_URL}/transaction`, transaction)
    .then(({ data }) => data);
}

export { fetchUserWalletDetails, fetchWalletBalance, transaction };
