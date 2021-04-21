package eth

import (
	"crypto/ecdsa"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"golang.org/x/crypto/pbkdf2"
)

const (
	AddressLength = 20
)

var secp256k1N, _ = new(big.Int).SetString("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", 16)

type EthereumExporter struct{}

func (exp *EthereumExporter) Name() string {
	return "Ethereum"
}

func (exp *EthereumExporter) ExportPrivateKey(key []byte) ([]byte, error) {
	curve := secp256k1.S256()

	// use pbkdf to extend the key
	if len(key) != curve.Params().BitSize/8 {
		keyLen := curve.Params().BitSize / 8
		key = pbkdf2.Key(key, []byte("ETH"), 1024, keyLen, sha1.New)
	}

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
	ret := fmt.Sprintf(
		`Account Address: %v
Private Key: %v`,
		crypto.PubkeyToAddress(priv.PublicKey),
		hex.EncodeToString(crypto.FromECDSA(&priv)))

	return []byte(ret), nil

}
