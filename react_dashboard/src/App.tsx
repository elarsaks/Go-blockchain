import React, { useState, useEffect } from "react";
import styled from "styled-components";
import AppHeader from "./components/AppHeader";
import Wallet from "./components/Wallet";
import BlockDiv from "./components/BlockDiv";
import { fetchBlockchainData } from "./api/Blockchain";
import { fetchWalletData } from "./api/Wallet";
import Notification from "./components/Notification";

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
  const [isError, setIsError] = useState({ message: "" });
  const [blockchain, setBlockchain] = useState<Block[]>([]);
  const [userWallet, setUserWallet] = useState<WalletContent>({
    blockchainAddress: "",
    privateKey: "",
    publicKey: "",
    amount: 0,
  });

  const fetchData = async () => {
    fetchBlockchainData()
      .then((blocks) => {
        setBlockchain(blocks);
        setIsLoading(false);
      })
      .catch((error) => {
        console.log(error);
        setIsError({ message: "Failed to fetch blockchain data" });
        setIsLoading(false);
      });

    fetchWalletData()
      .then((walletData) => setUserWallet(walletData))
      .catch((error) => {
        console.log(error);
        setIsError({ message: "Failed to fetch wallet data" });
        setIsLoading(false);
      });
  };

  useEffect(() => {
    fetchData();
  }, []);

  return (
    <div className="App">
      <AppHeader title="Go Blockchain" />
      <ContentContainer className="App">
        <WalletWrapperContainer>
          <Wallet walletContent={userWallet} />
          <Wallet walletContent={userWallet} />
        </WalletWrapperContainer>

        {isLoading && (
          <Notification type="info" message="Loading blockchain data." />
        )}

        {isError.message && (
          <Notification
            type="error"
            message="Sorry, there was an error loading blockchain data."
          />
        )}

        {!isLoading &&
          !isError.message &&
          blockchain.map((block, index) => (
            <React.Fragment key={index}>
              <BlockDiv block={block} />
            </React.Fragment>
          ))}
      </ContentContainer>
    </div>
  );
}

export default App;
