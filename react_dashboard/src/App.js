import React from "react";
import "./App.css";
import Header from "./components/Header/index";
// import EthOverview from "./components/Eth-Overview/index";
import Wallet from "./components/Wallet/index";

function App() {
  return (
    <div className="App">
      <Header />
      <div className="ui two column doubling stackable grid container">
        <div className="column">
          <Wallet />
        </div>

        <div className="column">
          <Wallet />
        </div>
      </div>

      {/* <EthOverview /> */}
    </div>
  );
}

export default App;
