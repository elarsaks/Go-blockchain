import axios from "axios";

/*
function snakeToCamelCase(snakeCaseString: string): string {
  return snakeCaseString.replace(/(_\w)/g, (match) => match[1].toUpperCase());
} */

function fetchWalletData(): Promise<WalletContent> {
  return (
    axios
      // TODO: Data type
      .post<any>("http://localhost:5000/wallet")
      .then(({ data }) => {
        const camelCaseResponseData: WalletContent = {
          blockchainAddress: data.blockchain_address,
          privateKey: data.private_key,
          publicKey: data.public_key,
          amount: 0, // TODO: Implement this
        };

        return camelCaseResponseData;
      })
      .catch((error) => {
        console.error("Failed to fetch wallet data:", error);
        return {
          blockchainAddress: "",
          privateKey: "",
          publicKey: "",
          amount: 0,
        };
      })
  );
}

export { fetchWalletData };
