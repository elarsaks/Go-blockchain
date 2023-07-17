import { fetchBlockchainData } from "./api/Blockchain";
import AppFooter from "./components/AppFooter";
import AppHeader from "./components/AppHeader";
import AppInfo from "./components/AppInfo";
import BlockDiv from "./components/BlockDiv";
import Loader from "./components/Loader";
import Notification from "./components/Notification";
import React, { useState, useEffect } from "react";
import styled from "styled-components";
import Wallet from "./components/Wallet";
const AppWrapper = styled.div`
  margin: 0;
`;

const ContentContainer = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  margin: 0;
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
        setBlockchain(blocks);
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
    <AppWrapper>
      <AppHeader title="Go Blockchain" />
      <ContentContainer className="App">
        <AppInfo />

        <WalletWrapperContainer>
          <Wallet type="miner" />
          <Wallet type="user" />
        </WalletWrapperContainer>

        {isLoading && (
          <Notification
            type="info"
            insideContainer={false}
            message="Loading blockchain data."
          />
        )}

        {isError.message && (
          <Notification
            type="error"
            message="Sorry, there was an error loading blockchain data."
            underDevelopment={true}
            insideContainer={false}
          />
        )}

        {!isLoading &&
          !isError.message &&
          blockchain.map((block, index) => (
            <React.Fragment key={index}>
              <Loader height={100} />
              <BlockDiv block={block} />
            </React.Fragment>
          ))}
      </ContentContainer>

      <AppFooter
        githubUrl="https://github.com/elarsaks"
        linkedinUrl="https://www.linkedin.com/in/elarsaks/"
        websiteUrl="https://saks.digital/"
      />
    </AppWrapper>
  );
}

export default App;
