package main

import (
	"github.com/xtaci/safebox/plugins/btc"
	"github.com/xtaci/safebox/plugins/dot"
	"github.com/xtaci/safebox/plugins/eth"
	"github.com/xtaci/safebox/plugins/fil"
	"github.com/xtaci/safebox/plugins/ssh"
	"github.com/xtaci/safebox/plugins/trx"
	"github.com/xtaci/safebox/plugins/xem"
)

type IKeyExport interface {
	Name() string                      // name of the exporter
	Desc() string                      // description
	Export(key []byte) ([]byte, error) // export function
	KeySize() int                      // expected key size to input to exporter
}

var exports []IKeyExport

func init() {
	exports = append(exports, new(eth.EthereumExporter))
	exports = append(exports, new(ssh.SSHExporter))
	exports = append(exports, new(btc.BitcoinExporter))
	exports = append(exports, new(fil.FileCoinExporter))
	exports = append(exports, new(trx.TronExporter))
	exports = append(exports, new(xem.NemExporter))
	exports = append(exports, new(dot.PolkadotExporter))
}
