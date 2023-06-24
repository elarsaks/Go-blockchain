import AppHeader from "./components/AppHeader";
import styled from "styled-components";
import Wallet from "./components/Wallet";
import BlockchainTable from "./components/BlockchainTable";
import dummyData from "./dummyData";

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

        <div>
          <h1>Blockchain Data</h1>
          <BlockchainTable blocks={dummyData} />
        </div>
      </ContentContainer>
    </div>
  );
}

export default App;
