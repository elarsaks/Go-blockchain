import React, { Dispatch, SetStateAction, useEffect, useState } from "react";
import styled from "styled-components";
import { fetchMinerWalletDetails } from "../../api/miner";
import { fetchUserWalletDetails, fetchWalletBalance } from "../../api/wallet";

const TitleRow = styled.div`
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 3rem;
`;

const MinerTitleContainer = styled.div`
  display: flex;
  align-items: center;
`;

const TypeSelect = styled.select`
  padding: 0.75rem 1.5rem;
  margin-right: 1rem;
  background-color: #ffffff;
  color: #00acd7;
  border: 1px solid #00acd7;
  border-radius: 5px;
  font-weight: bold;
  cursor: pointer;
`;

const Title = styled.h2`
  margin: 0 0 0 0;
`;

const Balance = styled.h2`
  margin: 0 0 0 0;
  color: #00acd7;
`;

interface WalletHeadProps {
  type: string;
  walletDetails: WalletState;
  setWalletDetails: Dispatch<SetStateAction<WalletState>>;
  setIsLoading: Dispatch<SetStateAction<boolean>>;
  setIsError: Dispatch<SetStateAction<LocalError>>;
}

const miners = [
  { value: "miner-1_1", text: "Miner 1" },
  { value: "miner-2", text: "Miner 2" },
  { value: "miner-3", text: "Miner 3" },
];

const WalletHead: React.FC<WalletHeadProps> = ({
  type,
  walletDetails,
  setWalletDetails,
  setIsLoading,
  setIsError,
}) => {
  const [selectedMiner, setSelectedMiner] = useState<{
    value: string;
    text: string;
  }>(miners[0]);

  const handleMinerChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    const selectedValue = event.target.value;
    const selectedMiner = miners.find((miner) => miner.value === selectedValue);

    if (selectedMiner) {
      setSelectedMiner(selectedMiner);
      fetchMinerDetails(selectedMiner.value);
    }
  };

  function fetchUserDetails() {
    setIsLoading(true);
    fetchUserWalletDetails()
      .then((userWalletDetails: WalletDetails) =>
        setWalletDetails((prevDetails) => ({
          ...prevDetails,
          ...userWalletDetails,
        }))
      )
      .catch((error: LocalError) => setIsError(error))
      .finally(() => setIsLoading(false));
  }

  function fetchMinerDetails(selectedMinerId: string) {
    setIsLoading(true);
    fetchMinerWalletDetails(selectedMinerId)
      .then((minerWalletDetails: WalletDetails) => {
        return fetchWalletBalance(minerWalletDetails.blockchainAddress).then(
          (balance) => {
            setWalletDetails((prevDetails) => ({
              ...prevDetails,
              ...minerWalletDetails,
              balance: balance === "0" ? "0.00" : balance,
            }));
            setIsError(null);
          }
        );
      })
      .catch((error: LocalError) => {
        setWalletDetails((prevDetails) => ({
          ...prevDetails,
          blockchainAddress: "",
          privateKey: "",
          publicKey: "",
          balance: "0.00",
        }));
        setIsError({ message: "Failed to fetch miner details" });
      })
      .finally(() => setIsLoading(false));
  }

  useEffect(() => {
    if (type === "User") {
      fetchUserDetails();
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [type]);

  useEffect(() => {
    if (type === "Miner") fetchMinerDetails("miner-1_1");
  }, [type, fetchMinerDetails]);

  useEffect(() => {
    let walletUpdate: NodeJS.Timeout;

    if (walletDetails.blockchainAddress) {
      walletUpdate = setInterval(() => {
        fetchWalletBalance(walletDetails.blockchainAddress)
          .then((balance) => {
            setWalletDetails((prevDetails) => ({
              ...prevDetails,
              balance: balance,
            }));
            setIsError(null);
          })
          .catch((error: LocalError) => setIsError(error));
      }, 10000);
    }

    return () => clearInterval(walletUpdate);
  }, [setIsError, setWalletDetails, walletDetails.blockchainAddress]);

  return (
    <div>
      {type === "User" ? (
        <TitleRow>
          <Title>User Wallet</Title>
          <Balance>{`${walletDetails.balance}₿`}</Balance>
        </TitleRow>
      ) : (
        <TitleRow>
          <MinerTitleContainer>
            <TypeSelect
              value={selectedMiner.value}
              onChange={handleMinerChange}
            >
              {miners.map((miner) => (
                <option key={miner.value} value={miner.value}>
                  {miner.text}
                </option>
              ))}
            </TypeSelect>
            <Title>{` Wallet`}</Title>
          </MinerTitleContainer>

          <Balance>{`${walletDetails.balance}₿`}</Balance>
        </TitleRow>
      )}
    </div>
  );
};

export default WalletHead;
