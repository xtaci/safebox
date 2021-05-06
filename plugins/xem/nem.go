package xem

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/isarq/nem-sdk-go/model"
	"github.com/xtaci/safebox/qrcode"
)

type NemExporter struct{}

func (exp *NemExporter) Name() string {
	return "Nem"
}

func (exp *NemExporter) Desc() string {
	return "NEM (New Economy Movement) is an ecosystem of platforms that use blockchain and cryptography to provide solutions for businesses and individuals. XEM is the native cryptocurrency of NEM’s NIS1 public blockchain."
}

func (exp *NemExporter) KeySize() int {
	return 32
}

func (exp *NemExporter) Export(key []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, errors.New("invalid key length")
	}

	keyPair := keyPairCreate(key)

	publicKey := keyPair.PublicString()
	privateKey := keyPair.PrivateString()

	address := model.ToAddress(publicKey, model.Data.Mainnet.ID)

	var out bytes.Buffer
	fmt.Fprintf(&out,
		`Nem Account: %v
Address QR Code:
%v
Private Key: %v
Private Key QR Code:
%v`,
		address,
		qrcode.GenerateQRCode(address),
		privateKey,
		qrcode.GenerateQRCode(privateKey),
	)

	return out.Bytes(), nil
}

func keyPairCreate(pk []byte) model.KeyPair {
	pair, err := model.FromSeed(pk)
	if err != nil {
		panic(err)
	}
	return pair
}
