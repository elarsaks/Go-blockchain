import { transaction } from "api/wallet";
import Notification from "components/shared/Notification";
import React, { useEffect, useState, useContext } from "react";
import styled from "styled-components";
import UtilReducer from "store/UtilReducer";
import WalletHead from "./WalletHead";
import { WalletContext } from "store/WalletProvider";

interface WalletContainerProps {
  isMiner: boolean;
}

const WalletContainer = styled.div<WalletContainerProps>`
  background-color: #f2f2f2;
  padding: 1rem;
  margin: 1rem;
  border: 1px solid #ccc;
  border-radius: 8px;
  width: 350px;

  @media (min-width: 850px) {
    margin-left: ${(props) => (props.isMiner ? "0" : "2rem")};
    margin-right: ${(props) => (props.isMiner ? "2rem" : "0")};
  }

  @media (max-width: 850px) {
    width: 80vw;
  }
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

interface ButtonProps {
  disabled: boolean;
}
const SendButton = styled.button<ButtonProps>`
  margin-top: 1rem;
  padding: 0.75rem 1.5rem;
  background-color: ${(props) => (props.disabled ? "#ccc" : "#00acd7")};
  color: white;
  border: none;
  border-radius: 3px;
  font-weight: bold;
  cursor: ${(props) => (props.disabled ? "not-allowed" : "pointer")};
  float: right;
  opacity: ${(props) => (props.disabled ? "0.6" : "1")};
`;

type WalletProps = {
  type: string;
};

const Wallet: React.FC<WalletProps> = ({ type }) => {
  const [isAnyFieldEmpty, setIsAnyFieldEmpty] = useState(false);

  const walletContext = useContext(WalletContext);

  const walletDetails =
    type === "Miner" ? walletContext.minerWallet : walletContext.userWallet;

  const setWalletDetails =
    type === "Miner"
      ? walletContext?.setMinerWallet
      : walletContext?.setUserWallet;

  const [utilState, dispatchUtil] = React.useReducer(UtilReducer, {
    isActive: false,
    type: "info",
    message: "",
  });

  useEffect(() => {
    setIsAnyFieldEmpty(
      walletDetails.blockchainAddress === "" ||
        walletDetails.privateKey === "" ||
        walletDetails.publicKey === "" ||
        walletDetails.recipientAddress === "" ||
        walletDetails.amount === ""
    );
  }, [walletDetails]);

  // Event Handlers
  const handleInputChange = (
    event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    const { name, value } = event.target;

    setWalletDetails({
      ...walletDetails,
      [name]: value,
    });
  };

  const sendCrypto = () => {
    transaction({
      message: "USER TRANSACTION",
      recipientBlockchainAddress: walletDetails.recipientAddress,
      senderBlockchainAddress: walletDetails.blockchainAddress,
      senderPrivateKey: walletDetails.privateKey,
      senderPublicKey: walletDetails.publicKey,
      value: walletDetails.amount,
    })
      .then((response) => {
        if (response.message === "fail") {
          dispatchUtil({
            type: "ON",
            payload: {
              type: "error",
              message: "Transaction failed",
            },
          });
        } else {
          dispatchUtil({
            type: "OFF",
            payload: null,
          });
        }
      })
      .catch((error) =>
        dispatchUtil({
          type: "ON",
          payload: {
            type: "error",
            message: error.message,
          },
        })
      );
  };

  return (
    <WalletContainer isMiner={type === "Miner"}>
      <WalletHead
        type={type}
        walletDetails={walletDetails}
        dispatchUtil={dispatchUtil}
      />

      <Form>
        <Field>
          <Label>Public Key</Label>
          <TextArea
            rows={4}
            name="publicKey"
            value={walletDetails.publicKey}
            onChange={handleInputChange}
          />
        </Field>

        <Field>
          <Label>Private Key</Label>
          <TextArea
            rows={2}
            name="privateKey"
            value={walletDetails.privateKey}
            onChange={handleInputChange}
          />
        </Field>

        <Field>
          <Label>
            {type === "Miner" ? "Miner " : "User"} Blockchain Address{" "}
          </Label>
          <TextArea
            rows={2}
            name="blockchainAddress"
            value={walletDetails.blockchainAddress}
            onChange={handleInputChange}
          />
        </Field>

        <Field>
          <Label>Recipient Blockchain Address</Label>
          <TextArea
            rows={2}
            name="recipientAddress"
            placeholder={
              type === "Miner"
                ? "User Blockchain Address"
                : "Miner Blockchain Address"
            }
            value={walletDetails.recipientAddress}
            onChange={handleInputChange}
          />
        </Field>

        <Field>
          <Label>Amount: </Label>
          <Input
            type="text"
            name="amount"
            placeholder="0.00₿"
            value={walletDetails.amount.toString()}
            onChange={handleInputChange}
          />
        </Field>

        {walletDetails.util.isActive && (
          <Notification
            type={walletDetails.util.type}
            message={walletDetails.util.message}
            underDevelopment={walletDetails.util.type !== "info"}
            insideContainer={false}
          />
        )}

        <SendButton
          type="submit"
          disabled={isAnyFieldEmpty}
          onClick={sendCrypto}
        >
          Send crypto
        </SendButton>
      </Form>
    </WalletContainer>
  );
};

export default Wallet;
