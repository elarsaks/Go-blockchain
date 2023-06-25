type Transaction = {
  sender_blockchain_address: string;
  recipient_blockchain_address: string;
  value: number;
};

type Block = {
  timestamp: number;
  nonce: number;
  previous_hash: string;
  transactions: Transaction[];
};

type Blockchain = { chain: Block[] };

export type { Blockchain, Block, Transaction };
