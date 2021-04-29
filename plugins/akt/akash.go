package akt

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/cosmos/go-bip39"
	"github.com/xtaci/safebox/qrcode"
	"strings"
)

type AkashExporter struct{}

func (exp *AkashExporter) Name() string {
	return "Akash"
}

func (exp *AkashExporter) KeySize() int {
	return 32
}

func (exp *AkashExporter) Export(key []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, errors.New("invalid key length")
	}

	mnemonic, err := bip39.NewMnemonic(key)
	if err != nil {
		return nil, err
	}
	//
	//configuration := types.GetConfig()
	//configuration.SetBech32PrefixForAccount(Bech32MainPrefix, Bech32MainPrefix+types.PrefixPublic)
	//
	//kb := keyring.NewInMemory()
	//hdPath := hd.NewFundraiserParams(0, types.GetConfig().GetCoinType(), 0).String()
	//account, err := kb.NewAccount(string(key), mnemonic, "", hdPath, hd.Secp256k1)
	//if err != nil {
	//	return nil, err
	//}
	//
	////Check Address
	//_, err = types.AccAddressFromBech32(account.GetAddress().String())
	//if err != nil {
	//	return nil, err
	//}
	mnemonicSlice := strings.Split(mnemonic, " ")
	var out bytes.Buffer
	fmt.Fprintf(&out,
		`Akash
Mnemonic: %v
Mnemonic QR Code(part1):
%v
Mnemonic QR Code(part2):
%v`,
		mnemonic,
		qrcode.GenerateQRCode(strings.Join(mnemonicSlice[0:len(mnemonicSlice)/2], " ")),
		qrcode.GenerateQRCode(strings.Join(mnemonicSlice[len(mnemonicSlice)/2:], " ")),
	)

	return out.Bytes(), nil
}

const (
	Bech32MainPrefix = "akash"
)
