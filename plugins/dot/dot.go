package dot

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/ChainSafe/gossamer/lib/common"
	"github.com/ChainSafe/gossamer/lib/crypto"
	"github.com/ChainSafe/gossamer/lib/crypto/sr25519"
	"github.com/btcsuite/btcutil/base58"
	"github.com/cosmos/go-bip39"
	"github.com/xtaci/safebox/qrcode"
	"golang.org/x/crypto/blake2b"
	"strings"
)

type PolkadotExporter struct{}

func (exp *PolkadotExporter) Name() string {
	return "Polkadot"
}

func (exp *PolkadotExporter) Desc() string {
	return "Polkadot is an open-source sharding multichain protocol that facilitates the cross-chain transfer of any data or asset types, not just tokens, thereby making a wide range of blockchains interoperable with each other."
}

func (exp *PolkadotExporter) KeySize() int {
	return 32
}

func (exp *PolkadotExporter) Export(key []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, errors.New("invalid key length")
	}

	mnemonic, err := bip39.NewMnemonic(key)
	if err != nil {
		return nil, err
	}
	kp, err := sr25519.NewKeypairFromMnenomic(mnemonic, "")
	if err != nil {
		return nil, err
	}

	addr, err := PublicKeyToAddress(kp.Public())
	if err != nil {
		return nil, err
	}
	mnemonicSlice := strings.Split(mnemonic, " ")
	var out bytes.Buffer
	fmt.Fprintf(&out,
		`Polkadot Account: %v
Address QR Code:
%v
Mnemonic: %v
Mnemonic QR Code(part1):
%v
Mnemonic QR Code(part2):
%v`,
		string(addr),
		qrcode.GenerateQRCode(string(addr)),

		mnemonic,
		qrcode.GenerateQRCode(strings.Join(mnemonicSlice[0:len(mnemonicSlice)/2], " ")),
		qrcode.GenerateQRCode(strings.Join(mnemonicSlice[len(mnemonicSlice)/2:], " ")),
	)

	return out.Bytes(), nil

}

var ss58Prefix = []byte("SS58PRE")

func PublicKeyToAddress(pub crypto.PublicKey) (common.Address, error) {
	enc := append([]byte{0}, pub.Encode()...)
	hasher, err := blake2b.New(64, nil)
	if err != nil {
		return "", err
	}
	_, err = hasher.Write(append(ss58Prefix, enc...))
	if err != nil {
		return "", err
	}
	checksum := hasher.Sum(nil)
	return common.Address(base58.Encode(append(enc, checksum[:2]...))), nil
}
