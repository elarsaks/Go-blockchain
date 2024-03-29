import React from "react";
import styled from "styled-components";

const BlockContainer = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  align-items: flex-start;
  background-color: #f2f2f2;
  padding: 1rem;
  margin: 1rem;
  border: 1px solid #ccc;
  border-radius: 8px;
  max-width: 800px;
  width: 100%;

  @media (max-width: 850px) {
    max-width: 80vw;
  }
`;

const Title = styled.h2`
  font-size: 1.5rem;
  text-align: left;
  color: #007acc;
  margin-bottom: 1rem;
  margin-top: 0;
`;

const TableTitle = styled.h3`
  color: #007acc;
  text-align: left;
  margin-bottom: 0;
`;

const Label = styled.span`
  font-weight: bold;
`;

const Value = styled.span`
  margin-left: 0.5rem;
  overflow-wrap: break-word;
  word-break: break-all;
`;

const NestedTable = styled.table`
  width: 100%;
  border-collapse: collapse;
  margin-top: 0.5rem;

  @media (max-width: 768px) {
    margin-top: 0.25rem;
  }
`;

const NestedTableHeader = styled.th`
  padding: 0.5rem;
  text-align: left;
  font-weight: bold;
  border-bottom: 1px solid #ccc;
`;

const NestedTableRow = styled.tr`
  &:nth-child(even) {
    background-color: #f2f2f2;
  }
`;

const NestedTableCell = styled.td`
  padding: 0.5rem;
  border-bottom: 1px solid #ccc;
  word-break: break-all;
`;

type BlockProps = {
  block: Block;
};

const BlockComponent: React.FC<BlockProps> = ({ block }) => (
  <BlockContainer>
    <Title>Block</Title>
    <div>
      <Label>Time:</Label>
      <Value>{block.timestamp}</Value>
    </div>
    <div>
      <Label>Nonce:</Label>
      <Value>{block.nonce}</Value>
    </div>
    <div>
      <Label>Previous hash:</Label>
      <Value>{block.previousHash}</Value>
    </div>

    <TableTitle>Transactions</TableTitle>
    <NestedTable>
      <thead>
        <tr>
          <NestedTableHeader>Sender</NestedTableHeader>
          <NestedTableHeader>Recipient</NestedTableHeader>
          <NestedTableHeader>Message</NestedTableHeader>
          <NestedTableHeader>Value</NestedTableHeader>
        </tr>
      </thead>
      <tbody>
        {block.transactions ? (
          block.transactions.map((transaction, idx) => (
            <NestedTableRow key={idx}>
              <NestedTableCell>
                {transaction.senderBlockchainAddress.length > 15
                  ? "..." + transaction.senderBlockchainAddress.slice(-15)
                  : transaction.senderBlockchainAddress}
              </NestedTableCell>
              <NestedTableCell>
                ...{transaction.recipientBlockchainAddress.slice(-15)}
              </NestedTableCell>
              <NestedTableCell>
                {transaction.message ? transaction.message : "_"}
              </NestedTableCell>
              <NestedTableCell>{transaction.value}</NestedTableCell>
            </NestedTableRow>
          ))
        ) : (
          <tr>
            <td colSpan={3}>No transactions (genesis block).</td>
          </tr>
        )}
      </tbody>
    </NestedTable>
  </BlockContainer>
);

export default BlockComponent;
