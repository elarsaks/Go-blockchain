import axios from "axios";

function fetchUserWalletDetails(): Promise<WalletDetails> {
  return (
    axios
      // TODO: Data type
      .post<any>("http://localhost:5000/wallet")
      .then(({ data }) => {
        const camelCaseResponseData: WalletDetails = {
          blockchainAddress: data.blockchain_address,
          privateKey: data.private_key,
          publicKey: data.public_key,
        };

        return camelCaseResponseData;
      })
  );
}

export { fetchUserWalletDetails };
