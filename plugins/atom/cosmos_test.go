package atom

import (
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/go-bip39"
	"testing"
)

func TestCosmosExporter_Export(t *testing.T) {
	key := [32]byte{}
	mnemonic, err := bip39.NewMnemonic(key[:])
	if err != nil {
		t.Log(err)
		return
	}
	t.Log("mnemonic:", mnemonic)
	account, err := NewAccount(mnemonic, types.FullFundraiserPath)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log("account:", account)
}
