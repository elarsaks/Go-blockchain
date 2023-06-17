import React from "react";
import "./wallet.css";

function Wallet() {
  return (
    <div class="ui form">
      <div class="field">
        <label>Text</label>
        <textarea></textarea>
      </div>
      <div class="field">
        <label>Short Text</label>
        <textarea rows="2"></textarea>
      </div>
    </div>
  );
}

export default Wallet;
