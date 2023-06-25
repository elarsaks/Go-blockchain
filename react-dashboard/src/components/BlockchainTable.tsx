import React, { useState } from "react";
import styled from "styled-components";

const StyledTable = styled.table`
  width: 100%;
  border-collapse: collapse;
  margin-top: 1rem;
`;

const TableHeader = styled.th`
  padding: 0.75rem;
  text-align: left;
  font-weight: bold;
  border-bottom: 1px solid #ccc;
`;

const TableRow = styled.tr`
  cursor: pointer;

  &:hover {
    background-color: #f2f2f2;
  }
`;

const TableCell = styled.td`
  padding: 0.75rem;
  border-bottom: 1px solid #ccc;
`;

const NestedTable = styled.table`
  width: 100%;
  border-collapse: collapse;

  border: 1px solid #00acd7;
  border-radius: 8px;
`;

const NestedTableHeader = styled.th`
  padding: 0.5rem;
  text-align: left;
  font-weight: bold;
  border-bottom: 1px solid #ccc;
  background-color: #f2f2f2;
`;

const NestedTableRow = styled.tr`
  background-color: #007acc;
`;

const NestedTableCell = styled.td`
  padding: 0.5rem;
  border-bottom: 1px solid #ccc;
`;

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

const BlockchainTable: React.FC<BlockchainTableProps> = ({ blocks }) => {
  const [expandedRow, setExpandedRow] = useState<number | null>(null);

  const handleRowClick = (index: number) => {
    setExpandedRow((prevRow) => (prevRow === index ? null : index));
  };

  const formatTimestamp = (timestamp: number): string => {
    const date = new Date(timestamp);
    return date.toLocaleString();
  };

  return (
    <StyledTable>
      <thead>
        <tr>
          <TableHeader>Timestamp</TableHeader>
          <TableHeader>Nonce</TableHeader>
          <TableHeader>Previous Hash</TableHeader>
        </tr>
      </thead>
      <tbody>
        {blocks.map((block, index) => (
          <React.Fragment key={index}>
            <TableRow onClick={() => handleRowClick(index)}>
              <TableCell>{formatTimestamp(block.timestamp)}</TableCell>
              <TableCell>{block.nonce}</TableCell>
              <TableCell>{block.previous_hash}</TableCell>
            </TableRow>
            {expandedRow === index && (
              <tr>
                <td colSpan={3}>
                  {block.transactions ? (
                    <NestedTable>
                      <thead>
                        <tr>
                          <NestedTableHeader>Sender</NestedTableHeader>
                          <NestedTableHeader>Recipient</NestedTableHeader>
                          <NestedTableHeader>Value</NestedTableHeader>
                        </tr>
                      </thead>
                      <tbody>
                        {block.transactions.map((transaction, idx) => (
                          <NestedTableRow key={idx}>
                            <NestedTableCell>
                              {transaction.sender_blockchain_address}
                            </NestedTableCell>
                            <NestedTableCell>
                              {transaction.recipient_blockchain_address}
                            </NestedTableCell>
                            <NestedTableCell>
                              {transaction.value}
                            </NestedTableCell>
                          </NestedTableRow>
                        ))}
                      </tbody>
                    </NestedTable>
                  ) : (
                    <p>No transactions</p>
                  )}
                </td>
              </tr>
            )}
          </React.Fragment>
        ))}
      </tbody>
    </StyledTable>
  );
};

export default BlockchainTable;
