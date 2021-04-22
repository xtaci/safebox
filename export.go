package main

import (
	"github.com/xtaci/safebox/eth"
	"github.com/xtaci/safebox/ssh"
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
}
