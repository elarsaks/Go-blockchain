import Header from "./components/Header/index";
import Wallet from "./components/Wallet/index";
import Table from "./components/Table/index";
import React from "react";
import "./App.css";

function App() {
  return (
    <div className="App">
      <Header />
      <br></br>

      <div className="ui two column doubling stackable grid container">
        <div className="column">
          <Wallet />
        </div>

        <div className="column">
          <Wallet />
        </div>
      </div>

      <br></br>
      <div className="ui container">
        <Table />
      </div>

      {/* <EthOverview /> */}
    </div>
  );
}

export default App;
