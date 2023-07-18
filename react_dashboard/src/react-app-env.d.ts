/// <reference types="react-scripts" />

type AmountResponse = {
  message: string;
  amount: string;
};

type Block = {
  timestamp: number;
  nonce: number;
  previousHash: string;
  transactions: Transaction[];
};

type Blockchain = {
  chain: Block[];
};

type LocalError = {
  message: string;
} | null;

type Transaction = {
  senderBlockchainAddress: string;
  recipientBlockchainAddress: string;
  value: string;
};

type WalletDetails = {
  blockchainAddress: string;
  privateKey: string;
  publicKey: string;
};

type WalletState = WalletDetails & {
  amount: string;
  recipientAddress: string;
};

type WalletDetailsResponse = {
  blockchain_address: string;
  private_key: string;
  public_key: string;
};
