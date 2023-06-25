import AppHeader from "./components/AppHeader";
import styled from "styled-components";
import Wallet from "./components/Wallet";
import BlockDiv from "./components/BlockDiv";
import dummyData from "./dummyData";
import React, { useState, useEffect } from "react";
import { fetchBlockchainData } from "./api/Blockchain";
import { Block } from "./Type"; // TODO: Fix types, they should not reuire import

const ContentContainer = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
`;

const WalletWrapperContainer = styled.div`
  display: flex;
  flex-wrap: wrap;
  gap: 1rem;
  justify-content: space-evenly;
`;

function App() {
  const [blockchain, setBlockchain] = useState<Block[]>([]);

  useEffect(() => {
    fetchData();
  }, []);

  const fetchData = async () => {
    try {
      const blockchainData = await fetchBlockchainData();
      setBlockchain(blockchainData);
      console.log(blockchainData); // Log the blockchain data
    } catch (error) {
      console.error("Failed to fetch blockchain data:", error);
    }
  };

  return (
    <div className="App">
      <AppHeader title="Go Blockchain" />
      <ContentContainer className="App">
        <WalletWrapperContainer>
          <Wallet />
          <Wallet />
        </WalletWrapperContainer>

        {dummyData.map((block, index) => (
          <React.Fragment key={index}>
            <BlockDiv block={block} />
          </React.Fragment>
        ))}
      </ContentContainer>
    </div>
  );
}

export default App;
