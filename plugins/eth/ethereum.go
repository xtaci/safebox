package eth

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"

	"github.com/btcsuite/btcd/btcec"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/xtaci/safebox/qrcode"
)

var secp256k1N, _ = new(big.Int).SetString("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", 16)

type EthereumExporter struct{}

func (exp *EthereumExporter) Name() string {
	return "Ethereum"
}

func (exp *EthereumExporter) Desc() string {
	return "Ethereum is a technology that's home to digital money, global payments, and applications."
}

func (exp *EthereumExporter) KeySize() int {
	return 32
}

func (exp *EthereumExporter) Export(key []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, errors.New("invalid key length")
	}
	curve := btcec.S256()

	var priv ecdsa.PrivateKey
	priv.Curve = curve
	priv.D = new(big.Int).SetBytes(key)
	priv.PublicKey.X, priv.PublicKey.Y = priv.PublicKey.Curve.ScalarBaseMult(priv.D.Bytes())
	if priv.PublicKey.X == nil {
		return nil, errors.New("invalid private key")
	}

	// The priv.D must < N
	if priv.D.Cmp(secp256k1N) >= 0 {
		return nil, fmt.Errorf("invalid private key, >=N")
	}
	// The priv.D must not be zero or negative.
	if priv.D.Sign() <= 0 {
		return nil, fmt.Errorf("invalid private key, zero or negative")
	}

	var out bytes.Buffer
	fmt.Fprintf(&out,
		`Ethereum Account: %v
Address QR Code:
%v
Private Key: %v
Private Key QR Code:
%v`,
		crypto.PubkeyToAddress(priv.PublicKey),
		qrcode.GenerateQRCode(crypto.PubkeyToAddress(priv.PublicKey).String()),

		hex.EncodeToString(crypto.FromECDSA(&priv)),
		qrcode.GenerateQRCode(hex.EncodeToString(crypto.FromECDSA(&priv))),
	)

	return out.Bytes(), nil

}
