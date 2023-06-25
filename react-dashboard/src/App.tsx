import AppHeader from "./components/AppHeader";
import styled from "styled-components";
import Wallet from "./components/Wallet";
import Block from "./components/Block";
import dummyData from "./dummyData";
import React, { useState } from "react";

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
            <Block block={block} />
          </React.Fragment>
        ))}
      </ContentContainer>
    </div>
  );
}

export default App;
