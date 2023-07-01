import React from "react";
import styled from "styled-components";

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

const Wallet: React.FC = () => (
  <WalletContainer>
    <Title>Wallet</Title>
    <Form>
      <Field>
        <Label>Public Key</Label>
        <TextArea rows={4} />
      </Field>
      <Field>
        <Label>Private Key</Label>
        <TextArea rows={2} />
      </Field>
      <Field>
        <Label>Sender Blockchain Address</Label>
        <TextArea rows={2} />
      </Field>
      <Field>
        <Label>Recipient Blockchain Address</Label>
        <TextArea rows={2} />
      </Field>
      <Field>
        <Label>Amount</Label>
        <Input type="text" placeholder="0" />
      </Field>
      <SubmitButton type="submit">Send crypto</SubmitButton>
    </Form>
  </WalletContainer>
);

export default Wallet;
