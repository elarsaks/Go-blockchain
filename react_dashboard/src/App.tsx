import React, { useState, useEffect } from "react";
import styled from "styled-components";
import AppHeader from "./components/AppHeader";
import Wallet from "./components/Wallet";
import BlockDiv from "./components/BlockDiv";
import { fetchBlockchainData } from "./api/Blockchain";
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

  function fetchchainData() {
    return fetchBlockchainData()
      .then((blocks) => {
        // setBlockchain(blocks);
        setIsError({ message: "Failed to fetch blockchain data" });
        setIsLoading(false);
      })
      .catch((error) => {
        setIsError({ message: "Failed to fetch blockchain data" });
        setIsLoading(false);
      });
  }

  useEffect(() => {
    // Fetch blockchain data immediately on component mount
    fetchchainData();

    // Fetch blockchain data every second
    const intervalId = setInterval(() => {
      fetchchainData();
    }, 1000);

    // Clean up function to clear the interval when the component unmounts
    return () => clearInterval(intervalId);
  }, []);

  return (
    <div className="App">
      <AppHeader title="Go Blockchain" />
      <ContentContainer className="App">
        <WalletWrapperContainer>
          <Wallet type="miner" />
          <Wallet type="user" />
        </WalletWrapperContainer>

        {isLoading && (
          <Notification type="info" message="Loading blockchain data." />
        )}

        {isError.message && (
          <Notification
            type="error"
            message="Sorry, there was an error loading blockchain data."
            underDevelopment={true}
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
