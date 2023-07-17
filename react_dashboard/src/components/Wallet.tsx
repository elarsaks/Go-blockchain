import styled from "styled-components";
import React, { useState, useEffect } from "react";
import {
  fetchMinerWalletDetails,
  fetchUserWalletDetails,
  fetchWalletAmount,
} from "../api/Wallet";
import Notification from "../components/Notification";

const WalletContainer = styled.div`
  background-color: #f2f2f2;
  padding: 1rem;
  margin: 1rem;
  border: 1px solid #ccc;
  border-radius: 8px;
  width: 350px;

  @media (max-width: 850px) {
    width: 100%;
  }
`;

const UserTitle = styled.h2`
  margin: 9px 0 24px 0;
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

const SendButton = styled.button`
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
  type: string;
};

const Wallet: React.FC<WalletProps> = ({ type }) => {
  // State
  const [isLoading, setIsLoading] = useState(true);
  const [isError, setIsError] = useState<LocalError>(null);

  const [selectedMiner, setSelectedMiner] = useState({
    value: "miner1",
    text: "Miner 1",
    url:
      process.env.MINER_1_WALLET_ADDRESS ||
      "http://localhost:5001/miner/wallet",
  });

  const [miners] = useState([
    selectedMiner,
    {
      value: "miner2",
      text: "Miner 2",
      url:
        process.env.MINER_2_WALLET_ADDRESS ||
        "http://localhost:5002/miner/wallet",
    },
    {
      value: "miner3",
      text: "Miner 3",
      url:
        process.env.MINER_3_WALLET_ADDRESS ||
        "http://localhost:5003/miner/wallet",
    },
  ]);

  const [walletDetails, setWalletDetails] = useState<WalletDetails>({
    blockchainAddress: "",
    privateKey: "",
    publicKey: "",
  });

  const [walletAmount, setWalletAmount] = useState(0);

  // Methods
  function fetchAmount(blockchainAddress: string) {
    fetchWalletAmount(blockchainAddress)
      .then((walletAmount) => setWalletAmount(walletAmount))
      .catch((error: LocalError) =>
        setError({ message: "Failed to fetch wallet amount" })
      );
  }

  function fetchMinerDetails(selectedMinerUrl: string) {
    setIsLoading(true);
    fetchMinerWalletDetails(selectedMinerUrl)
      .then((walletDetails: WalletDetails) => {
        fetchAmount(walletDetails.blockchainAddress);
        setWalletDetails(walletDetails);
        setIsLoading(false);
      })
      .catch((error: LocalError) =>
        setError({ message: "Failed to fetch MINER details" })
      );
  }

  function fetchUserDetails() {
    fetchUserWalletDetails()
      .then((walletDetails: WalletDetails) => {
        fetchAmount(walletDetails.blockchainAddress);
        setWalletDetails(walletDetails);
        setIsLoading(false);
      })
      .catch((error: LocalError) =>
        setError({ message: "Failed to fetch USER details" })
      );
  }

  function setError(error: LocalError) {
    setIsError(error);
    setIsLoading(false);
  }

  const handleMinerChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    const selectedValue = event.target.value;

    const selectedMinerObj = miners.find(
      (miner) => miner.value === selectedValue
    );

    setSelectedMiner(selectedMinerObj ? selectedMinerObj : selectedMiner);
    fetchMinerDetails(selectedMiner.url);
  };

  // Effects
  useEffect(() => {
    if (type === "user") fetchUserDetails();
  }, [type]);

  useEffect(() => {
    if (type === "miner") {
      fetchMinerDetails(selectedMiner.url);
    }
  }, [type, selectedMiner.url]);

  useEffect(() => {
    let walletUpdate: NodeJS.Timeout;

    if (walletDetails.blockchainAddress) {
      walletUpdate = setInterval(() => {
        fetchAmount(walletDetails.blockchainAddress);
      }, 3000);
    }

    return () => clearInterval(walletUpdate);
  }, [walletDetails.blockchainAddress]);

  // Event Handlers
  const handleInputChange = (
    event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    const { name, value } = event.target;

    setWalletDetails((prevDetails) => ({
      ...prevDetails, // spread to keep the existing state
      [name]: value, // update the value of the property defined by 'name'
    }));
  };

  const handleSubmit = (event: React.FormEvent) => {
    event.preventDefault();
    // Handle form submission if needed
    // You can access the updated wallet content in localWalletContent state
  };

  return (
    <WalletContainer>
      {type === "user" ? (
        <UserTitle>User Wallet</UserTitle>
      ) : (
        <TitleRow>
          <MinerTitle>{`${selectedMiner.text} Wallet`}</MinerTitle>
          <TypeSelect value={selectedMiner.value} onChange={handleMinerChange}>
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
            {type === "miner" ? selectedMiner.text : "User"} Blockchain Address{" "}
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
          <TextArea rows={2} />
        </Field>

        <Field>
          <Label>Amount</Label>
          <Input
            type="text"
            name="amount"
            placeholder="0"
            value={walletAmount.toString()}
            onChange={handleInputChange}
          />
        </Field>

        {isLoading && (
          <Notification
            type="info"
            message="Loading data."
            underDevelopment={false}
            insideContainer={true}
          />
        )}

        {isError && (
          <Notification
            type="error"
            message={isError.message || "Something went wrong."}
            underDevelopment={true}
            insideContainer={true}
          />
        )}
        <SendButton type="submit" disabled={isError !== null}>
          Send crypto
        </SendButton>
      </Form>
    </WalletContainer>
  );
};

export default Wallet;
