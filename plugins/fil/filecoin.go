package fil

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"errors"
	"fmt"
	"io"

	secp256k1 "github.com/ethereum/go-ethereum/crypto"
	"github.com/filecoin-project/go-address"
	"github.com/xtaci/safebox/qrcode"
)

type FileCoinExporter struct{}

func (exp *FileCoinExporter) Name() string {
	return "FileCoin"
}
func (exp *FileCoinExporter) Desc() string {
	return "Filecoin is a decentralized storage network designed to store humanity's most important information."
}
func (exp *FileCoinExporter) KeySize() int {
	return 64
}

func (exp *FileCoinExporter) Export(key []byte) ([]byte, error) {
	if len(key) != 64 {
		return nil, errors.New("invalid key length")
	}
	address.CurrentNetwork = address.Mainnet
	priv, err := GenerateKeyFromSeed(bytes.NewBuffer(key))
	if err != nil {
		return nil, err
	}

	ki := KeyInfo{
		Type:       KTSecp256k1,
		PrivateKey: priv,
	}
	keys, err := NewKey(ki)
	if err != nil {
		return nil, err
	}

	var out bytes.Buffer
	fmt.Fprintf(&out,
		`FileCoin Account: %v
Address QR Code:
%v
Private Key: %v
Private Key QR Code :
%v`,
		keys.Address.String(),
		qrcode.GenerateQRCode(keys.Address.String()),
		hex.EncodeToString(keys.PrivateKey),
		qrcode.GenerateQRCode(hex.EncodeToString(keys.PrivateKey)),
	)

	return out.Bytes(), nil
}

// KeyType defines a type of a key
type KeyType string

const (
	KTSecp256k1     KeyType = "secp256k1"
	PrivateKeyBytes         = 32
)

// KeyInfo is used for storing keys in KeyStore
type KeyInfo struct {
	Type       KeyType
	PrivateKey []byte
}

func NewKey(keyinfo KeyInfo) (*Key, error) {
	k := &Key{
		KeyInfo: keyinfo,
	}

	var err error
	k.PublicKey = PublicKey(k.PrivateKey)
	switch k.Type {
	case KTSecp256k1:
		k.Address, err = address.NewSecp256k1Address(k.PublicKey)
		if err != nil {
			return nil, fmt.Errorf("converting Secp256k1 to address: %w", err)
		}
	default:
		return nil, fmt.Errorf("unsupported key type: %s", k.Type)
	}
	return k, nil
}

type Key struct {
	KeyInfo
	PublicKey []byte
	Address   address.Address
}

func GenerateKeyFromSeed(seed io.Reader) ([]byte, error) {
	key, err := ecdsa.GenerateKey(secp256k1.S256(), seed)
	if err != nil {
		return nil, err
	}

	privkey := make([]byte, PrivateKeyBytes)
	blob := key.D.Bytes()

	// the length is guaranteed to be fixed, given the serialization rules for secp2561k curve points.
	copy(privkey[PrivateKeyBytes-len(blob):], blob)

	return privkey, nil
}

func PublicKey(sk []byte) []byte {
	x, y := secp256k1.S256().ScalarBaseMult(sk)
	return elliptic.Marshal(secp256k1.S256(), x, y)
}
