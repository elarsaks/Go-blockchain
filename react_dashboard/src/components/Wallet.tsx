import styled from "styled-components";
import React, { useState, useEffect } from "react";

const WalletContainer = styled.div`
  background-color: #f2f2f2;
  padding: 1rem;
  margin: 1rem;
  border: 1px solid #ccc;
  border-radius: 8px;
  width: 350px;
`;

const UserTitle = styled.h2`
  margin: 0;
`;

const TitleRow = styled.div`
  display: flex;
  justify-content: space-between;
  align-items: center;
`;

const MinerTitle = styled.h2`
  margin: 0;
`;

const TypeSelect = styled.select`
  padding: 0.75rem 1.5rem;
  background-color: #ffffff;
  color: #00acd7;
  border: 1px solid #00acd7;
  border-radius: 5px;
  font-weight: bold;
  cursor: pointer;
`;

const Form = styled.div`
  margin-top: 1rem;
`;

const Field = styled.div`
  margin-bottom: 1rem;
`;

const Label = styled.label`
  display: block;
  margin-bottom: 0.5rem;
  font-weight: bold;
  text-align: left;
`;

const TextArea = styled.textarea`
  width: 95%;
  padding: 0.5rem;
  text-align: left;
`;

const Input = styled.input`
  width: 95%;
  padding: 0.5rem;
  text-align: left;
`;

const SubmitButton = styled.button`
  margin-top: 1rem;
  padding: 0.75rem 1.5rem;
  background-color: #00acd7;
  color: white;
  border: none;
  border-radius: 3px;
  font-weight: bold;
  cursor: pointer;
  float: right;
`;

type WalletProps = {
  walletContent: WalletContent;
  type: string;
};

const Wallet: React.FC<WalletProps> = ({ walletContent, type }) => {
  const [localWalletContent, setLocalWalletContent] = useState(walletContent);
  const [selectedMiner, setSelectedMiner] = useState("miner1");
  const [miners, setMiners] = useState([
    { value: "miner1", text: "Miner 1" },
    { value: "miner2", text: "Miner 2" },
    { value: "miner3", text: "Miner 3" },
  ]);

  const selectedMinerText =
    miners.find((miner) => miner.value === selectedMiner)?.text || "";

  useEffect(() => {
    setLocalWalletContent(walletContent);
  }, [walletContent]);

  const handleMinerChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    setSelectedMiner(event.target.value);
  };

  const handleInputChange = (
    event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    const { name, value } = event.target;

    setLocalWalletContent((prevState: WalletContent) => ({
      ...prevState,
      [name]: value,
    }));
  };
  const handleSubmit = (event: React.FormEvent) => {
    event.preventDefault();
    // Handle form submission if needed
    // You can access the updated wallet content in localWalletContent state
  };

  return (
    <WalletContainer>
      {type === "User" ? (
        <UserTitle>User Wallet</UserTitle>
      ) : (
        <TitleRow>
          <MinerTitle>{`${selectedMinerText} Wallet`}</MinerTitle>
          <TypeSelect value={selectedMiner} onChange={handleMinerChange}>
            {miners.map((miner) => (
              <option key={miner.value} value={miner.value}>
                {miner.text}
              </option>
            ))}
          </TypeSelect>
        </TitleRow>
      )}

      <Form onSubmit={handleSubmit}>
        <Field>
          <Label>Public Key</Label>
          <TextArea
            rows={4}
            name="publicKey"
            value={localWalletContent.publicKey}
            onChange={handleInputChange}
          />
        </Field>
        <Field>
          <Label>Private Key</Label>
          <TextArea
            rows={2}
            name="privateKey"
            value={localWalletContent.privateKey}
            onChange={handleInputChange}
          />
        </Field>
        <Field>
          <Label>Sender Blockchain Address</Label>
          <TextArea
            rows={2}
            name="blockchainAddress"
            value={localWalletContent.blockchainAddress}
            onChange={handleInputChange}
          />
        </Field>
        <Field>
          <Label>Recipient Blockchain Address</Label>
          <TextArea rows={2} />
        </Field>
        <Field>
          <Label>Amount</Label>
          <Input
            type="text"
            name="amount"
            placeholder="0"
            value={localWalletContent.amount.toString()}
            onChange={handleInputChange}
          />
        </Field>
        <SubmitButton type="submit">Send crypto</SubmitButton>
      </Form>
    </WalletContainer>
  );
};

export default Wallet;
