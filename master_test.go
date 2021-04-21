package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testMKeyPath = "./testkey"
	testPass     = "passwd"
)

func TestCreateKey(t *testing.T) {
	mkey := newMasterKey()
	mkey.generateMasterKey(nil)
	mkey.lables[0] = "HELLO"
	assert.Nil(t, mkey.store([]byte(testPass), testMKeyPath))

	loaded := newMasterKey()
	assert.Nil(t, loaded.load([]byte(testPass), testMKeyPath))

	assert.Equal(t, mkey.lables[0], loaded.lables[0])
	os.Remove(testMKeyPath)
}
