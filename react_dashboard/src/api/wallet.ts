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

function fetchWalletBalance(blockchainAddress: string): Promise<string> {
  return axios
    .get<AmountResponse>(
      `http://localhost:5000/wallet/amount?blockchain_address=${blockchainAddress}`
    )
    .then(({ data }) => data.amount);
}

function transaction(transaction: Transaction): Promise<string> {
  console.log(transaction);
  // Why this string ends up in golang as a number is beyond me
  return axios
    .post<string>(`http://localhost:5000/transaction`, transaction)
    .then(({ data }) => data);
}

export {
  fetchMinerWalletDetails,
  fetchUserWalletDetails,
  fetchWalletBalance,
  transaction,
};
