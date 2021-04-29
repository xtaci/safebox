package band

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/cosmos/go-bip39"
	"github.com/xtaci/safebox/qrcode"
	"strings"
)

type BandExporter struct{}

func (exp *BandExporter) Name() string {
	return "Band"
}

func (exp *BandExporter) KeySize() int {
	return 32
}

func (exp *BandExporter) Export(key []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, errors.New("invalid key length")
	}

	mnemonic, err := bip39.NewMnemonic(key)
	if err != nil {
		return nil, err
	}

	//configuration := types.GetConfig()
	//configuration.SetBech32PrefixForAccount(Bech32MainPrefix, Bech32MainPrefix+types.PrefixPublic)
	//configuration.SetCoinType(Bip44CoinType)
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
		`Band
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
	Bech32MainPrefix = "band"
	Bip44CoinType    = 494
)
