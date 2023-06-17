import React, { Component } from "react";
import axios from "axios";
import "./eth-overview.css";
import { Card, Grid, Icon } from "semantic-ui-react";
import LatestBlocks from "../Latest-Blocks/index";
import LatestTxs from "../Latest-Txs/index";

// import api key from the env variable
const apiKey = process.env.REACT_APP_ETHERSCAN_API_KEY;

const endpoint = `https://api.etherscan.io/api`;

class EthOverview extends Component {
  constructor() {
    super();
    this.state = {
      ethUSD: "",
      ethBTC: "",
      blockNo: "",
      latestBlock: 0,
      difficulty: "",
      marketCap: 0,
    };
  }

  async componentDidMount() {
    // get the ethereum price
    const prices = await axios.get(
      endpoint + `?module=stats&action=ethprice&apikey=${apiKey}`
    );
    let { result } = prices.data;
    this.setState({
      ethUSD: result.ethusd,
      ethBTC: result.ethbtc,
    });

    // get the market cap of ether in USD
    const marketCap = await axios.get(
      endpoint + `?module=stats&action=ethsupply&apikey=${apiKey}`
    );

    result = marketCap.data.result;
    // in wei
    const priceWei = result.toString();

    // in ether
    const priceEth = priceWei.slice(0, priceWei.length - 18);
    console.log(result, priceWei, priceEth);
    // convert eth in USD
    this.setState({
      marketCap: parseInt(priceEth) * this.state.ethUSD,
    });

    // get the latest block number
    const latestBlock = await axios.get(
      endpoint + `?module=proxy&action=eth_blockNumber&apikey=${apiKey}`
    );
    this.setState({
      latestBlock: parseInt(latestBlock.data.result),
      blockNo: latestBlock.data.result, // save block no in hex
    });

    // get the block difficulty
    const blockDetail = await axios.get(
      endpoint +
        `?module=proxy&action=eth_getBlockByNumber&tag=${latestBlock.data.result}&boolean=true&apikey=${apiKey}`
    );
    result = blockDetail.data.result;

    const difficulty = parseInt(result.difficulty).toString();

    // convert difficulty in Terra Hash
    // instead of dividing it with 10^12 we'll slice it
    const difficultyTH = `${difficulty.slice(0, 4)}.${difficulty.slice(
      4,
      6
    )} TH`;

    this.setState({
      difficulty: difficultyTH,
    });
  }

  getLatestBlocks = () => {
    if (this.state.latestBlock) {
      return <LatestBlocks latestBlock={this.state.latestBlock}></LatestBlocks>;
    }
  };

  getLatestTxs = () => {
    if (this.state.blockNo) {
      return <LatestTxs blockNo={this.state.blockNo}></LatestTxs>;
    }
  };

  render() {
    const { ethUSD, ethBTC, latestBlock, difficulty, marketCap } = this.state;
    return (
      <div>
        <Grid>
          <Grid.Row>
            <Grid.Column width={4}>
              <Card>
                <Card.Content>
                  <Card.Header style={{ color: "#1d6fa5" }}>
                    <Icon name="ethereum"></Icon> ETHER PRICE
                  </Card.Header>
                  <Card.Description textAlign="left">
                    <Icon name="usd"></Icon>
                    {ethUSD} <Icon name="at"></Icon> {ethBTC}{" "}
                    <Icon name="bitcoin"></Icon>
                  </Card.Description>
                </Card.Content>
              </Card>
            </Grid.Column>
            <Grid.Column width={4}>
              <Card>
                <Card.Content>
                  <Card.Header style={{ color: "#1d6fa5" }}>
                    <Icon name="list alternate outline"></Icon> LATEST BLOCK
                  </Card.Header>
                  <Card.Description textAlign="left">
                    <Icon name="square"></Icon> {latestBlock}
                  </Card.Description>
                </Card.Content>
              </Card>
            </Grid.Column>
            <Grid.Column width={4}>
              <Card>
                <Card.Content>
                  <Card.Header style={{ color: "#1d6fa5" }}>
                    <Icon name="setting"></Icon> DIFFICULTY
                  </Card.Header>
                  <Card.Description textAlign="left">
                    {difficulty}
                  </Card.Description>
                </Card.Content>
              </Card>
            </Grid.Column>
            <Grid.Column width={4}>
              <Card>
                <Card.Content>
                  <Card.Header style={{ color: "#1d6fa5" }}>
                    <Icon name="world"></Icon> MARKET CAP
                  </Card.Header>
                  <Card.Description textAlign="left">
                    <Icon name="usd"></Icon> {marketCap}
                  </Card.Description>
                </Card.Content>
              </Card>
            </Grid.Column>
          </Grid.Row>
        </Grid>

        <Grid divided="vertically">
          <Grid.Row columns={2}>
            <Grid.Column>{this.getLatestBlocks()}</Grid.Column>
            <Grid.Column>{this.getLatestTxs()}</Grid.Column>
          </Grid.Row>
        </Grid>
      </div>
    );
  }
}

export default EthOverview;
