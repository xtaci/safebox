package trx

import (
	"bytes"
	"crypto/ecdsa"
	"errors"
	"fmt"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/sasaxie/go-client-api/common/base58"
	"github.com/sasaxie/go-client-api/common/crypto"
	"github.com/xtaci/safebox/qrcode"
)

type TronExporter struct{}

func (exp *TronExporter) Name() string {
	return "Tron"
}

func (exp *TronExporter) Desc() string {
	return "Tron is a blockchain-based decentralized platform that aims to build a free, global digital content entertainment system with distributed storage technology, and allows easy and cost-effective sharing of digital content."
}

func (exp *TronExporter) KeySize() int {
	return 64
}

func (exp *TronExporter) Export(key []byte) ([]byte, error) {
	if len(key) != 64 {
		return nil, errors.New("invalid key length")
	}

	k, err := ecdsa.GenerateKey(ethcrypto.S256(), bytes.NewBuffer(key))
	if err != nil {
		return nil, err
	}

	priv := k.D.Text(16)
	address := crypto.PubkeyToAddress(k.PublicKey)
	addressString := base58.EncodeCheck(address.Bytes())

	var out bytes.Buffer
	fmt.Fprintf(&out,
		`Tron Account: %v
Address QR Code:
%v
Private Key: %v
Private Key QR Code:
%v`,
		addressString,
		qrcode.GenerateQRCode(addressString),
		priv,
		qrcode.GenerateQRCode(priv),
	)

	return out.Bytes(), nil
}
