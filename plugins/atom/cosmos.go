package atom

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/go-bip39"
	"github.com/xtaci/safebox/qrcode"
	"strings"
)

type CosmosExporter struct{}

func (exp *CosmosExporter) Name() string {
	return "Cosmos"
}

func (exp *CosmosExporter) KeySize() int {
	return 32
}

func (exp *CosmosExporter) Export(key []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, errors.New("invalid key length")
	}

	mnemonic, err := bip39.NewMnemonic(key)
	if err != nil {
		return nil, err
	}

	kb := keyring.NewInMemory()
	hdPath := types.FullFundraiserPath
	account, err := kb.NewAccount(string(key), mnemonic, "", hdPath, hd.Secp256k1)
	if err != nil {
		return nil, err
	}
	//Check Address
	_, err = types.AccAddressFromBech32(account.GetAddress().String())
	if err != nil {
		return nil, err
	}
	mnemonicSlice := strings.Split(mnemonic, " ")
	var out bytes.Buffer
	fmt.Fprintf(&out,
		`Cosmos Account: %v
HD derivation path: %v
Public Key QR Code:
%v
Mnemonic: %v
Mnemonic QR Code(part1):
%v
Mnemonic QR Code(part2):
%v`,
		account.GetAddress().String(),
		hdPath,
		qrcode.GenerateQRCode(account.GetPubKey().String()),
		mnemonic,
		qrcode.GenerateQRCode(strings.Join(mnemonicSlice[0:len(mnemonicSlice)/2], " ")),
		qrcode.GenerateQRCode(strings.Join(mnemonicSlice[len(mnemonicSlice)/2:], " ")),
	)

	return out.Bytes(), nil
}
