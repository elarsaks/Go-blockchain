/// <reference types="react-scripts" />

type Transaction = {
  senderBlockchainAddress: string;
  recipientBlockchainAddress: string;
  value: number;
};

type Block = {
  timestamp: number;
  nonce: number;
  previousHash: string;
  transactions: Transaction[];
};

type WalletContent = {
  blockchainAddress: string;
  privateKey: string;
  publicKey: string;
  amount: number;
};

type Blockchain = {
  chain: Block[];
};
