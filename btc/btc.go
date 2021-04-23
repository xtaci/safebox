package btc

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/xtaci/safebox/qrcode"
	"golang.org/x/crypto/pbkdf2"
)

const (
	AddressLength = 20
)

var secp256k1N, _ = new(big.Int).SetString("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", 16)

type BitcoinExporter struct{}

func (exp *BitcoinExporter) Name() string {
	return "Bitcoin"
}
func (exp *BitcoinExporter) KeySize() int {
	return 32
}

func (exp *BitcoinExporter) Export(key []byte) ([]byte, error) {
	curve := btcec.S256()
	// use pbkdf to extend the key
	if len(key) != curve.Params().BitSize/8 {
		keyLen := curve.Params().BitSize / 8
		key = pbkdf2.Key(key, []byte("ETH"), 1024, keyLen, sha1.New)
	}
	priv, pub := btcec.PrivKeyFromBytes(curve, key)

	address, err := btcutil.NewAddressPubKey(pub.SerializeUncompressed(), &chaincfg.MainNetParams)
	if err != nil {
		return nil, err
	}

	var out bytes.Buffer
	fmt.Fprintf(&out,
		`Bitcoin Account: %v
Public Key QR Code:
%v
Private Key: %v
Private Key QR Code :
%v`,
		address.EncodeAddress(),
		qrcode.GenerateQRCode(address.EncodeAddress()),

		hex.EncodeToString(priv.Serialize()),
		qrcode.GenerateQRCode(string(priv.Serialize())),
	)

	return out.Bytes(), nil

}
