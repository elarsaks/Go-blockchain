import { fetchBlockchainData } from "api/miner";
import AppFooter from "components/layout/AppFooter";
import AppHeader from "components/layout/AppHeader";
import AppInfo from "components/layout/AppInfo";
import Background from "components/layout/Background";
import BlockDiv from "components/BlockDiv";
import Loader from "components/shared/Loader";
import Notification from "components/shared/Notification";
import React, { useEffect, useState, useCallback } from "react";
import styled from "styled-components";
import UtilReducer from "store/UtilReducer";
import Wallet from "components/wallet/Wallet";
import { WalletProvider } from "store/WalletProvider";

const AppWrapper = styled.div`
  margin: 0;
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  overflow: auto;
`;

const ContentContainer = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
`;

const WalletWrapperContainer = styled.div`
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;

  @media (max-width: 850px) {
    justify-content: center;
  }
`;

function App() {
  const [blockchain, setBlockchain] = useState<Block[]>([]);
  const [utilState, dispatchUtil] = React.useReducer(UtilReducer, {
    isActive: false,
    type: "info",
    message: "",
  });

  const fetchchainData = useCallback(() => {
    if (blockchain.length === 0) {
      dispatchUtil({
        type: "ON",
        payload: {
          type: "info",
          message: "Fetching blockchain data...",
        },
      });
    }

    return fetchBlockchainData()
      .then((blocks) => {
        setBlockchain(blocks);
        dispatchUtil({
          type: "OFF",
          payload: null,
        });
      })
      .catch((error) => {
        dispatchUtil({
          type: "ON",
          payload: {
            type: "error",
            message: error.message,
          },
        });
      });
  }, [blockchain.length]);

  useEffect(() => {
    // Fetch blockchain data immediately on component mount
    fetchchainData();

    // Fetch blockchain data every 3 seconds
    const intervalId = setInterval(() => {
      fetchchainData();
    }, 5000);

    // Clear interval on component unmount
    return () => clearInterval(intervalId);
  }, [fetchchainData]);

  return (
    <AppWrapper>
      <Background />
      <AppHeader title="Go Blockchain" />
      <ContentContainer className="App">
        <AppInfo />

        <WalletWrapperContainer>
          <WalletProvider previousHash={blockchain[0]?.previousHash}>
            <Wallet type="Miner" />
            <Wallet type="User" />
          </WalletProvider>
        </WalletWrapperContainer>

        {utilState.isActive && (
          <Notification
            type={utilState.type}
            message={utilState.message}
            underDevelopment={true}
            insideContainer={false}
          />
        )}

        {!utilState.isActive &&
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
