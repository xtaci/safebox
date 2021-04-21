package main

import (
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
	assert.Nil(t, mkey.save([]byte(testPass), testMKeyPath))

	loaded := newMasterKey()
	assert.Nil(t, loaded.load([]byte(testPass), testMKeyPath))
}
