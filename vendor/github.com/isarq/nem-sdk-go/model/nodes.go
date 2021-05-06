package model

type SearchTestnet struct {
	Url, Location string
}

type NetNode struct {
	Uri string
}

// The default testnet node
//const DefaultTestnet = `http://bigalice2.nem.ninja`
const DefaultTestnet = `http://192.3.61.243`

// The default mainnet node
const DefaultMainnet = `http://alice6.nem.ninja`

// The default mijin node
const DefaultMijin = ``

// The default mainnet block explorer
const MainnetExplorer = `http://chain.nem.ninja/#/transfer/`

// The default testnet block explorer
const TestnetExplorer = `http://bob.nem.ninja:8765/#/transfer/`

// The default mijin block explorer
const MijinExplorer = ``

// The nodes allowing search by transaction hash on testnet
// type slice	Location
var SearchOnTestnet = []SearchTestnet{
	{
		Url:      `http://bigalice2.nem.ninja`,
		Location: `America / New_York`,
	}, {
		Url:      `http://192.3.61.243`,
		Location: `America / Los_Angeles`,
	}, {
		Url:      `http://23.228.67.85`,
		Location: `America / Los_Angeles`,
	},
}

// The nodes allowing search by transaction hash on mainnet
// type slice
var SearchOnMainnet = []SearchTestnet{
	{
		Url:      `http: //62.75.171.41`,
		Location: `Germany`,
	}, {
		Url:      `http: //104.251.212.131`,
		Location: `USA`,
	}, {
		Url:      `http: //45.124.65.125`,
		Location: `Hong Kong`,
	}, {
		Url:      `http: //185.53.131.101`,
		Location: `Netherlands`,
	}, {
		Url:      `http: //sz.nemchina.com`,
		Location: `China`,
	},
}

// The nodes allowing search by transaction hash on mijin
// type slice
var SearchOnMijin = []SearchTestnet{
	{
		Url:      ``,
		Location: ``,
	},
}

// The testnet nodes
// type slice
var TestnetNode = []NetNode{
	{Uri: `http://104.128.226.60`},
	{Uri: `http: //23.228.67.85`},
	{Uri: `http: //192.3.61.243`},
	{Uri: `http: //50.3.87.123`},
	{Uri: `http: //localhost`},
}

// The mainnet nodes
// type slice
var MainnetNode = []NetNode{
	{Uri: `http: //62.75.171.41`},
	{Uri: `ttp: //san.nem.ninja`},
	{Uri: `http: //go.nem.ninja`},
	{Uri: `http: //hachi.nem.ninja`},
	{Uri: `http: //jusan.nem.ninja`},
	{Uri: `http: //nijuichi.nem.ninja`},
	{Uri: `http: //alice2.nem.ninja`},
	{Uri: `http: //alice3.nem.ninja`},
	{Uri: `http: //alice4.nem.ninja`},
	{Uri: `http: //alice5.nem.ninja`},
	{Uri: `http: //alice6.nem.ninja`},
	{Uri: `http: //alice7.nem.ninja`},
	{Uri: `http: //localhost`},
}

// The mijin nodes
// type slice
var MijinNode = []NetNode{
	{
		Uri: ``,
	},
}

// The server verifying signed apostilles
const ApostilleAuditServer = `http://185.117.22.58:4567/verify`

// The API to get all supernodes
const Supernodes = `https://supernodes.nem.io/nodes`

// The API to get XEM/BTC market data
const MarketInfo = `https://poloniex.com/public`

//The API to get BTC/USD market data
const BtcPrice = `https://blockchain.info/ticker`

// The default endpoint port
const DefaultPort = 7890

// The Mijin endpoint port
const mijinPort = 7895

// The websocket port
const WebsocketPort = 7778
