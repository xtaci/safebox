package atom

import (
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
	//kb := keyring.NewInMemory()
	//account, err := kb.NewAccount("string(key)", mnemonic, "", types.FullFundraiserPath, hd.Secp256k1)
	//if err!=nil{
	//	t.Log(err)
	//	return
	//}
	//t.Log("account:",account)
}
