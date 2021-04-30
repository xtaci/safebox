package fil

import (
	"bytes"
	"crypto/rand"
	"io"
	"testing"
)

func TestFileCoinExporter_Export(t *testing.T) {
	key := make([]byte, 64)
	io.ReadFull(rand.Reader, key)
	priv, err := GenerateKeyFromSeed(bytes.NewBuffer(key))
	if err != nil {
		t.Error("err:", err)
		return
	}
	ki := KeyInfo{
		Type:       KTSecp256k1,
		PrivateKey: priv,
	}
	keys, err := NewKey(ki)
	if err != nil {
		t.Error("err:", err)
		return
	}
	t.Log("address:", keys.Address.String())
}
