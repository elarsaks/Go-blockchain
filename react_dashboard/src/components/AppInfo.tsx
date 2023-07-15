import React from "react";
import styled from "styled-components";

const AppInfoWrapper = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: flex-start;
  padding: 1em;
  margin: 1em;
  border-radius: 5px;
  color: #ffffff;
  width: 90%;
  max-width: 785px;
  overflow: auto;
  background-color: #00add8;
  border: 1px solid #007d9c;
`;

const AppInfo: React.FC = () => (
  <AppInfoWrapper>
    <h3>This is a simple example of a blockchain.</h3>
    <p>
      The wallet on the left represents a miner, while the wallet on the right
      represents a random user. Miner wallets accumulate crypto when they mine
      blocks.
    </p>
    <p>
      You can experiment by sending this crypto from miners to users and vice
      versa. Just copy and paste the wallet address from one wallet to the other
      wallets recipient input.
    </p>
    <p>Beneath the wallets, you'll find the 10 most recently mined blocks.</p>
  </AppInfoWrapper>
);

export default AppInfo;
