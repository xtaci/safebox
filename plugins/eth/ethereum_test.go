package eth

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExport(t *testing.T) {
	key := make([]byte, 32)
	io.ReadFull(rand.Reader, key)
	t.Log("key:", hex.EncodeToString(key))

	eth := new(EthereumExporter)
	priv, err := eth.Export(key)
	assert.Nil(t, err)
	t.Log("output:", string(priv))
}
