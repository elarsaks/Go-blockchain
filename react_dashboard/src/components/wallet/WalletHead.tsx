import React, { Dispatch, SetStateAction, useEffect, useState } from "react";
import styled from "styled-components";
import {
  fetchMinerWalletDetails,
  fetchUserWalletDetails,
  fetchWalletAmount,
  transaction,
} from "../../api/Wallet";

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

interface WalletHeadProps {
  type: string;
  walletDetails: WalletState;
  setWalletDetails: Dispatch<SetStateAction<WalletState>>;
  setIsLoading: Dispatch<SetStateAction<boolean>>;
  setIsError: Dispatch<SetStateAction<LocalError>>;
}

const WalletHead: React.FC<WalletHeadProps> = (props) => {
  const selectedMinerUrls = {
    miner1: process.env.REACT_APP_MINER_1 + "/miner/wallet",
    miner2: process.env.REACT_APP_MINER_2 + "/miner/wallet",
    miner3: process.env.REACT_APP_MINER_3 + "/miner/wallet",
  };

  const miners = [
    { value: "miner1", text: "Miner 1", url: selectedMinerUrls.miner1 },
    { value: "miner2", text: "Miner 2", url: selectedMinerUrls.miner2 },
    { value: "miner3", text: "Miner 3", url: selectedMinerUrls.miner3 },
  ];

  const [selectedMiner, setSelectedMiner] = useState<{
    value: string;
    text: string;
    url: string;
  }>({
    value: "miner1",
    text: "Miner 1",
    url: process.env.REACT_APP_MINER_1 + "/miner/wallet",
  });

  const handleMinerChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    const selectedValue = event.target.value;
    const selectedMiner = miners.find((miner) => miner.value === selectedValue);

    if (selectedMiner) {
      setSelectedMiner(selectedMiner);
      fetchMinerDetails(selectedMiner.url);
    }
  };

  function fetchWalletDetails(walletDetails: WalletDetails) {
    // setIsLoading(true);
    fetchWalletAmount(walletDetails.blockchainAddress);
    /*  .then((balance) =>
        setWalletDetails((prevDetails) => ({
          ...prevDetails,
          ...walletDetails,
          amount: balance,
        }))
      ) 
      .catch((error: LocalError) =>
        setError({ message: "Failed to fetch wallet details" })
      )
      .finally(() => setIsLoading(false)); */
  }

  function fetchUserDetails() {
    fetchUserWalletDetails();
    /*  .then((userWalletDetails: WalletDetails) =>
        fetchWalletDetails(userWalletDetails)
      )
      .catch((error: LocalError) =>
        setError({ message: "Failed to fetch user details" })
      ); */
  }

  function fetchMinerDetails(selectedMinerUrl: string) {
    fetchMinerWalletDetails(selectedMinerUrl);
    /*  .then((minerWalletDetails: WalletDetails) =>
        fetchWalletDetails(minerWalletDetails)
      )
      .catch((error: LocalError) =>
        setError({ message: "Failed to fetch miner details" })
      ); */
  }

  // Effects
  // TODO: Fix using effects whitout disabling eslint (Learn React Hooks)
  useEffect(() => {
    if (props.type === "user") {
      fetchUserDetails();
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [props.type]);

  useEffect(() => {
    if (props.type === "miner") {
      fetchMinerDetails(selectedMiner.url);
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [props.type, selectedMiner.url]);

  useEffect(() => {
    let walletUpdate: NodeJS.Timeout;

    if (props.walletDetails.blockchainAddress) {
      walletUpdate = setInterval(() => {
        fetchWalletAmount(props.walletDetails.blockchainAddress);
        /*  .then((balance) =>
            setWalletDetails((prevDetails) => ({
              ...prevDetails,
              amount: balance,
            }))
          )
          .catch((error: LocalError) =>
            setError({ message: "Failed to fetch wallet amount" })
          ); */
      }, 3000);
    }

    return () => clearInterval(walletUpdate);
  }, [props.walletDetails.blockchainAddress]);

  return (
    <div>
      {props.type === "user" ? (
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
    </div>
  );
};

export default WalletHead;
