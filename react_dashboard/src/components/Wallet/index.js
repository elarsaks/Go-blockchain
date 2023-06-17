import React from "react";
import "./wallet.css";

function Wallet() {
  return (
    <div className="ui segment ">
      <h2 className="ui header">Wallet</h2>
      <div className="ui form">
        <div className="field">
          <label>Public Key</label>
          <textarea rows="4"></textarea>
        </div>

        <div className="field">
          <label>Private Key</label>
          <textarea rows="2"></textarea>
        </div>

        <div className="field">
          <label>Sender Blockchain Address</label>
          <textarea rows="2"></textarea>
        </div>

        <div className="field">
          <label>Recipient Blockchain Address</label>
          <textarea rows="2"></textarea>
        </div>

        <div className="field">
          <label>Amount</label>
          <input type="text" placeholder="0"></input>
        </div>

        <button className="ui button" type="submit">
          Send crypto
        </button>
      </div>
    </div>
  );
}

export default Wallet;
