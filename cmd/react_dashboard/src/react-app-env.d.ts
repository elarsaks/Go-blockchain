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
  message: string;
  recipientBlockchainAddress: string;
  senderBlockchainAddress: string;
  senderPrivateKey: string;
  senderPublicKey: string;
  value: string;
};

type MiningContextType = {
  mining: boolean;
  setMining: React.Dispatch<React.SetStateAction<boolean>>;
};

type UtilState = {
  isActive: boolean;
  type: "info" | "warning" | "error" | "success";
  message: string;
};

// If type is "ON", then payload is an object with type and message properties.
// If type is "OFF", then payload is null.
type UtilAction =
  | {
      type: "ON";
      payload: {
        type: "info" | "warning" | "error" | "success";
        message: string;
      };
    }
  | {
      type: "OFF";
      payload: null;
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
  blockchainAddress: string;
  privateKey: string;
  publicKey: string;
};

type WalletStore = {
  minerWallet: WalletState;
  userWallet: WalletState;
  setMinerWallet: React.Dispatch<React.SetStateAction<WalletState>>;
  setUserWallet: React.Dispatch<React.SetStateAction<WalletState>>;
};
