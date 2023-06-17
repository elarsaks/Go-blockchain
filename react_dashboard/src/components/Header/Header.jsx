import React from "react";
// import Header component from the semantic-ui-react
import { Header } from "semantic-ui-react";
import "./header.css";

function AppDashboard() {
  return (
    <div>
      <Header as="h2" block>
        React | Golang | Blockchain
      </Header>
    </div>
  );
}

export default AppDashboard;
