package btc

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/xtaci/safebox/qrcode"
)

const (
	AddressLength = 20
)

var secp256k1N, _ = new(big.Int).SetString("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", 16)

type BitcoinExporter struct{}

func (exp *BitcoinExporter) Name() string {
	return "Bitcoin"
}
func (exp *BitcoinExporter) Desc() string {
	return "Bitcoin is an innovative payment network and a new kind of money."
}
func (exp *BitcoinExporter) KeySize() int {
	return 32
}

func (exp *BitcoinExporter) Export(key []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, errors.New("invalid key length")
	}
	priv, pub := btcec.PrivKeyFromBytes(key)

	address, err := btcutil.NewAddressPubKey(pub.SerializeUncompressed(), &chaincfg.MainNetParams)
	if err != nil {
		return nil, err
	}

	var out bytes.Buffer
	fmt.Fprintf(&out,
		`Bitcoin Account: %v
Address QR Code:
%v
Private Key: %v
Private Key QR Code:
%v`,
		address.EncodeAddress(),
		qrcode.GenerateQRCode(address.EncodeAddress()),

		hex.EncodeToString(priv.Serialize()),
		qrcode.GenerateQRCode(string(priv.Serialize())),
	)

	return out.Bytes(), nil

}
