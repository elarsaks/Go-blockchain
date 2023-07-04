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

const Title = styled.h2`
  font-size: 1.5rem;
  text-align: left;
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
};

const Wallet: React.FC<WalletProps> = ({ walletContent }) => {
  const [localWalletContent, setLocalWalletContent] = useState(walletContent);

  useEffect(() => {
    setLocalWalletContent(walletContent);
  }, [walletContent]);

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
      <Title>Wallet</Title>
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
