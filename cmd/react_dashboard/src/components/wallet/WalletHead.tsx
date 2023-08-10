import { WalletContext } from "store/WalletProvider";
import React, { Dispatch, useContext, useState } from "react";
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
      //   fetchMinerDetails(selectedMiner.value);
    }
  };

  const walletContext = useContext(WalletContext);

  const walletDetails =
    type === "Miner" ? walletContext.minerWallet : walletContext.userWallet;

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
