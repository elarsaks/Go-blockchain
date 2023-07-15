import React from "react";
import styled from "styled-components";

const AppInfoWrapper = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: flex-start;
  padding: 1em;
  margin: 0;
  margin-top: 1em;
  width: 90%;
  max-width: 785px;
  border-radius: 5px;
  color: #ffffff;
  overflow: auto;
  background-color: #00add8;
  border: 1px solid #007d9c;

  h3 {
    margin-bottom: 5px;
    margin-top: 5px;
  }

  p {
    margin-bottom: 5px;
    margin-top: 5px;
  }

  a {
    font-size: 18px; 
    color: black; 
    background-color: 
    border: none; 
    text-decoration: none; 
  }
`;

const AppInfo: React.FC = () => (
  <AppInfoWrapper>
    <h3> This is a simple example of a blockchain.</h3>
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
    <br />
    <p>
      <b>
        Check out the project on{" "}
        <a
          href="https://github.com/elarsaks/Go-blockchain"
          target="_blank"
          rel="noopener noreferrer"
        >
          GitHub
        </a>{" "}
      </b>
      .
    </p>
  </AppInfoWrapper>
);

export default AppInfo;
