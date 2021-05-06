package atom

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec/legacy"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/go-bip39"
	"github.com/xtaci/safebox/qrcode"
)

type CosmosExporter struct{}

func (exp *CosmosExporter) Name() string {
	return "Cosmos"
}

func (exp *CosmosExporter) Desc() string {
	return "Cosmos is an ever-expanding ecosystem of interoperable and sovereign blockchain apps and services, built for a decentralized future"
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

	hdPath := types.FullFundraiserPath
	account, err := NewAccount(mnemonic, hdPath)
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

func NewAccount(mnemonic string, hdPath string) (*LocalInfo, error) {
	// create master key and derive first key for keyring
	derivedPriv, err := hd.Secp256k1.Derive()(mnemonic, "", hdPath)
	if err != nil {
		return nil, err
	}
	privKey := hd.Secp256k1.Generate()(derivedPriv)
	pub := privKey.PubKey()
	info := newLocalInfo(pub, string(legacy.Cdc.MustMarshalBinaryBare(privKey)), hd.Secp256k1.Name())
	return info, nil
}

type LocalInfo struct {
	PubKey       cryptotypes.PubKey `json:"pubkey"`
	PrivKeyArmor string             `json:"privkey.armor"`
	Algo         hd.PubKeyType      `json:"algo"`
}

func newLocalInfo(pub cryptotypes.PubKey, privArmor string, algo hd.PubKeyType) *LocalInfo {
	return &LocalInfo{
		PubKey:       pub,
		PrivKeyArmor: privArmor,
		Algo:         algo,
	}
}

func (i LocalInfo) GetPubKey() cryptotypes.PubKey {
	return i.PubKey
}

func (i LocalInfo) GetAddress() types.AccAddress {
	return i.PubKey.Address().Bytes()
}

func (i LocalInfo) GetAlgo() hd.PubKeyType {
	return i.Algo
}
