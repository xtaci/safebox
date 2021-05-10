package main

import (
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
)

const (
	testMKeyPath = "./testkey"
	testPass     = "passwd"
)

func TestCreateKey(t *testing.T) {
	mkey := newMasterKey()
	mkey.changePassword([]byte(testPass))
	mkey.generateMasterKey(nil)
	mkey.labels[0] = "HELLO"
	assert.Nil(t, mkey.store(testMKeyPath))

	loaded := newMasterKey()
	assert.Nil(t, loaded.load([]byte(testPass), testMKeyPath))

	assert.Equal(t, mkey.labels[0], loaded.labels[0])
	os.Remove(testMKeyPath)
	t.Log(hexutil.Encode(mkey.masterKey[:100]))
}

func TestCreateKeySalted(t *testing.T) {
	mkey := newMasterKey()
	mkey.generateMasterKey([]byte("added some salt"))
	mkey.changePassword([]byte(testPass))
	mkey.labels[0] = "HELLO"
	assert.Nil(t, mkey.store(testMKeyPath))

	loaded := newMasterKey()
	assert.Nil(t, loaded.load([]byte(testPass), testMKeyPath))

	assert.Equal(t, mkey.labels[0], loaded.labels[0])
	os.Remove(testMKeyPath)

	t.Log(hexutil.Encode(mkey.masterKey[:100]))
}
