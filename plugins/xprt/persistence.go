package xprt

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/go-bip39"
	"github.com/xtaci/safebox/plugins/atom"
	"github.com/xtaci/safebox/qrcode"
)

type PersistenceExporter struct{}

func (exp *PersistenceExporter) Name() string {
	return "Persistence"
}

func (exp *PersistenceExporter) Desc() string {
	return "Persistence is an interoperable protocol built to facilitate the creation of next-gen financial products."
}

func (exp *PersistenceExporter) KeySize() int {
	return 32
}

func (exp *PersistenceExporter) Export(key []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, errors.New("invalid key length")
	}

	mnemonic, err := bip39.NewMnemonic(key)
	if err != nil {
		return nil, err
	}

	configuration := types.GetConfig()
	configuration.SetBech32PrefixForAccount(Bech32MainPrefix, Bech32MainPrefix+types.PrefixPublic)
	configuration.SetCoinType(Bip44CoinType)

	hdPath := hd.NewFundraiserParams(0, types.GetConfig().GetCoinType(), 0).String()
	account, err := atom.NewAccount(mnemonic, hdPath)
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
		`Persistence Account: %v
HD derivation path: %v
Address QR Code:
%v
Mnemonic: %v
Mnemonic QR Code(part1):
%v
Mnemonic QR Code(part2):
%v`,
		account.GetAddress().String(),
		hdPath,
		qrcode.GenerateQRCode(account.GetAddress().String()),
		mnemonic,
		qrcode.GenerateQRCode(strings.Join(mnemonicSlice[0:len(mnemonicSlice)/2], " ")),
		qrcode.GenerateQRCode(strings.Join(mnemonicSlice[len(mnemonicSlice)/2:], " ")),
	)

	return out.Bytes(), nil
}

const (
	Bech32MainPrefix = "persistence"
	Bip44CoinType    = 750
)
