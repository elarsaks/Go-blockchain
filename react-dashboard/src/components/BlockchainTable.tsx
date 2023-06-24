import React from "react";

type Transaction = {
  sender_blockchain_address: string;
  recipient_blockchain_address: string;
  value: number;
};

type Block = {
  timestamp: number;
  nonce: number;
  previous_hash: string;
  transactions: Transaction[] | null;
};

type BlockchainTableProps = {
  blocks: Block[];
};

const BlockchainTable: React.FC<BlockchainTableProps> = ({ blocks }) => (
  <table>
    <thead>
      <tr>
        <th>Timestamp</th>
        <th>Nonce</th>
        <th>Previous Hash</th>
        <th>Transactions</th>
      </tr>
    </thead>
    <tbody>
      {blocks.map((block, index) => (
        <tr key={index}>
          <td>{block.timestamp}</td>
          <td>{block.nonce}</td>
          <td>{block.previous_hash}</td>
          <td>
            {block.transactions ? (
              <ul>
                {block.transactions.map((transaction, idx) => (
                  <li key={idx}>
                    <strong>Sender:</strong>{" "}
                    {transaction.sender_blockchain_address},{" "}
                    <strong>Recipient:</strong>{" "}
                    {transaction.recipient_blockchain_address},{" "}
                    <strong>Value:</strong> {transaction.value}
                  </li>
                ))}
              </ul>
            ) : (
              "None"
            )}
          </td>
        </tr>
      ))}
    </tbody>
  </table>
);

export default BlockchainTable;
