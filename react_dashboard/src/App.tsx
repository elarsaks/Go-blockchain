import AppHeader from "./components/AppHeader";
import styled from "styled-components";
import Wallet from "./components/Wallet";
import BlockDiv from "./components/BlockDiv";
import React, { useState, useEffect } from "react";
import { fetchBlockchainData } from "./api/Blockchain";
import { fetchWalletData } from "./api/Wallet";
import { Block, WalletContent /* Blockchain */ } from "./Type";

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
  const [isLoading, setIsLoading] = useState(true);
  const [blockchain, setBlockchain] = useState<Block[]>([]);
  const [userWallet, setUserWallet] = useState<WalletContent>({
    blockchainAddress: "",
    privateKey: "",
    publicKey: "",
    amount: 0,
  });

  useEffect(() => {
    fetchData();
  }, []);

  const fetchData = async () => {
    try {
      const blockchainData = await fetchBlockchainData();
      setBlockchain(blockchainData.chain);

      const userWalletData = await fetchWalletData();
      setUserWallet(userWalletData);

      setIsLoading(false);
    } catch (error) {
      // TODO: Handle error
      console.error("Failed to fetch blockchain data:", error);
      // setIsLoading(false);
    }
  };

  useEffect(() => {
    if (blockchain != null && blockchain.length > 0) {
      console.log("Blockchain: ", blockchain);
      setIsLoading(false);
    }
  }, [blockchain]);

  return (
    <div className="App">
      <AppHeader title="Go Blockchain" />
      <ContentContainer className="App">
        <WalletWrapperContainer>
          <Wallet walletContent={userWallet} />
          <Wallet walletContent={userWallet} />
        </WalletWrapperContainer>

        {isLoading ? (
          <p>Loading blockchain data... </p>
        ) : (
          blockchain.map((block, index) => (
            <React.Fragment key={index}>
              <BlockDiv block={block} />
            </React.Fragment>
          ))
        )}
      </ContentContainer>
    </div>
  );
}

export default App;
