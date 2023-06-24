import React from "react";
import logo from "./logo.svg";
import "./App.css";
import { Button } from "./components/Button";

function App() {
  return (
    <div className="App">
      <Button>Normal</Button>
      <Button $primary>Primary</Button>
    </div>
  );
}

export default App;
