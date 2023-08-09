import { fetchMinerWalletDetails } from "api/miner";
import { fetchUserWalletDetails, fetchWalletBalance } from "api/wallet";
import { WalletContext } from "store/WalletProvider";
import React, { Dispatch, useContext, useEffect, useState } from "react";
import styled from "styled-components";

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

const TypeSelect = styled.select<{ disabled?: boolean }>`
  padding: 0.75rem 1.5rem;
  margin-right: 1rem;
  background-color: ${(props) => (props.disabled ? "#f0f0f0" : "#ffffff")};
  color: ${(props) => (props.disabled ? "#a0a0a0" : "#00acd7")};
  border: 1px solid ${(props) => (props.disabled ? "#a0a0a0" : "#00acd7")};
  border-radius: 5px;
  font-weight: bold;
  cursor: ${(props) => (props.disabled ? "not-allowed" : "pointer")};
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
  dispatchUtil: Dispatch<UtilAction>;
}

const miners = [
  { value: "1", text: "Miner 1" },
  { value: "2", text: "Miner 2" },
  { value: "3", text: "Miner 3" },
];

const WalletHead: React.FC<WalletHeadProps> = ({ type, dispatchUtil }) => {
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

  const walletContext = useContext(WalletContext);

  const walletDetails =
    type === "Miner" ? walletContext.minerWallet : walletContext.userWallet;

  const setWalletDetails =
    type === "Miner"
      ? walletContext.setMinerWallet
      : walletContext.setUserWallet;

  function fetchUserDetails() {
    dispatchUtil({
      type: "ON",
      payload: {
        type: "info",
        message: "User wallet will be regitered when next block is mined...",
      },
    });

    fetchUserWalletDetails()
      .then((userWalletDetails: WalletDetails) => {
        setWalletDetails((prevDetails) => ({
          ...prevDetails,
          ...userWalletDetails,
        }));
        dispatchUtil({
          type: "OFF",
          payload: null,
        });
      })
      .catch((error: LocalError) =>
        dispatchUtil({
          type: "ON",
          payload: {
            type: "error",
            message: "Failed to fetch user wallet details",
          },
        })
      );
  }

  function fetchMinerDetails(selectedMinerId: string) {
    dispatchUtil({
      type: "ON",
      payload: {
        type: "info",
        message: "Fetching miner wallet details...",
      },
    });

    // Fetch miner wallet details
    return (
      fetchMinerWalletDetails(selectedMinerId)
        .then((minerWalletDetails: WalletDetails) => {
          setWalletDetails((prevDetails) => ({
            ...prevDetails,
            ...minerWalletDetails,
          }));

          dispatchUtil({
            type: "OFF",
            payload: null,
          });

          return minerWalletDetails.blockchainAddress;
        })

        // Fetch miner wallet balance
        .then((blockchainAddress) =>
          fetchWalletBalance(blockchainAddress).then((balance) => {
            setWalletDetails((prevDetails) => ({
              ...prevDetails,
              balance: balance === "0" ? "0.00" : balance,
            }));
          })
        )
    );
  }

  useEffect(() => {
    if (type === "User") {
      fetchUserDetails();
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [type]);

  useEffect(
    () => {
      if (type === "Miner") fetchMinerDetails(selectedMiner.value);
    },
    // eslint-disable-next-line react-hooks/exhaustive-deps
    [type, selectedMiner.value]
  );

  useEffect(() => {
    let walletUpdate: NodeJS.Timeout;

    if (walletDetails.blockchainAddress) {
      walletUpdate = setInterval(() => {
        fetchWalletBalance(walletDetails.blockchainAddress).then((balance) => {
          setWalletDetails((prevDetails) => ({
            ...prevDetails,
            balance: balance === "0" ? "0.00" : balance,
          }));
        });
      }, 10000);
    }

    return () => clearInterval(walletUpdate);
  }, [setWalletDetails, walletDetails.blockchainAddress]);

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
              disabled={true}
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
