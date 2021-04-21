package main

import "github.com/xtaci/safebox/eth"

type IKeyExport interface {
	Name() string
	Export(key []byte) ([]byte, error)
	KeySize() int
}

var exports []IKeyExport

func init() {
	exports = append(exports, new(eth.EthereumExporter))
}
