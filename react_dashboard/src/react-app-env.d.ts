/// <reference types="react-scripts" />

type BalanceResponse = {
  error: string;
  balance: string;
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
  recipientBlockchainAddress: string;
  senderBlockchainAddress: string;
  senderPrivateKey: string;
  senderPublicKey: string;
  value: string;
};

type WalletDetails = {
  blockchainAddress: string;
  privateKey: string;
  publicKey: string;
};

type WalletState = WalletDetails & {
  amount: string;
  balance: string;
  recipientAddress: string;
};

type WalletDetailsResponse = {
  blockchain_address: string;
  private_key: string;
  public_key: string;
};
