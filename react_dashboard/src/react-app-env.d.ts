/// <reference types="react-scripts" />

type AmountResponse = {
  message: string;
  amount: number;
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
  value: number;
};

type WalletDetails = {
  blockchainAddress: string;
  privateKey: string;
  publicKey: string;
};

type WalletDetailsResponse = {
  blockchain_address: string;
  private_key: string;
  public_key: string;
};
