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
  return axios.post<WalletDetailsResponse>(minerAdress).then(({ data }) => {
    const camelCaseResponseData: WalletDetails = {
      blockchainAddress: data.blockchain_address,
      privateKey: data.private_key,
      publicKey: data.public_key,
    };

    return camelCaseResponseData;
  });
}

function fetchWalletAmount(blockchainAddress: string): Promise<number> {
  return axios
    .get<AmountResponse>(
      `http://localhost:5000/wallet/amount?blockchain_address=${blockchainAddress}`
    )
    .then(({ data }) => data.amount);
}

function transaction(
  senderAddress: string,
  receiverAddress: string,
  amount: number
): Promise<string> {
  return axios
    .post<string>(`http://localhost:5000/transaction`, {
      senderBlockchainAddress: senderAddress,
      recipientBlockchainAddress: receiverAddress,
      value: amount,
    })
    .then(({ data }) => data);
}

export {
  fetchMinerWalletDetails,
  fetchUserWalletDetails,
  fetchWalletAmount,
  transaction,
};
