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

const TableContainer = styled.div`
  display: flex;
  background-color: #f2f2f2;
  padding: 1.9rem;
  margin-top: 1rem;
  border: 1px solid #ccc;
  border-radius: 8px;
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

        <TableContainer>
          <BlockchainTable blocks={dummyData} />
        </TableContainer>
      </ContentContainer>
    </div>
  );
}

export default App;
