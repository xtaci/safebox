package main

import (
	"github.com/xtaci/safebox/plugins/btc"
	"github.com/xtaci/safebox/plugins/eth"
	"github.com/xtaci/safebox/plugins/ssh"
)

type IKeyExport interface {
	Name() string
	Export(key []byte) ([]byte, error)
	KeySize() int
}

var exports []IKeyExport

func init() {
	exports = append(exports, new(eth.EthereumExporter))
	exports = append(exports, new(ssh.SSHExporter))
	exports = append(exports, new(btc.BitcoinExporter))
}
