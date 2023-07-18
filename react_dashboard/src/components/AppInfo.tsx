import styled from "styled-components";
import React, { useState, useEffect } from "react";

const AppInfoWrapper = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: flex-start;
  padding: 1em;
  margin: 0;
  margin-top: 1em;
  width: 90%;
  max-width: 800px;
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

  @media (max-width: 850px) {
    width: 80vw;
  }
`;

const AppInfo: React.FC = () => {
  const [isMobile, setIsMobile] = useState(false);

  useEffect(() => {
    const handleResize = () => {
      setIsMobile(window.innerWidth <= 850);
    };

    handleResize(); // Check initial width
    window.addEventListener("resize", handleResize);

    return () => {
      window.removeEventListener("resize", handleResize);
    };
  }, []);

  return (
    <AppInfoWrapper>
      <h3> This is a simple example of a blockchain.</h3>
      <p>
        The wallet on the {isMobile ? "up" : "left"} represents a miner, while
        the wallet on the {isMobile ? "down" : "right"} represents a random
        user. Miner wallets accumulate crypto when they mine blocks. At first,
        they mine 10 blocks, and then they mine 1 block only if there are
        transactions to be mined.
      </p>
      <p>
        You can experiment by sending this crypto from miners to users and vice
        versa. Just copy and paste the wallet address from one wallet to the
        other wallets recipient input.
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
};

export default AppInfo;
