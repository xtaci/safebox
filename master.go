package main

import (
	"crypto/rand"
	"io"
)

const MasterKeyLength = 256 * 1024 // 256 KB
type MasterKey struct {
	masterKey [MasterKeyLength]byte
}

type DerivedKey [16]byte

func newMasterKey() *MasterKey {
	mkey := new(MasterKey)
	return mkey
}

func (mkey *MasterKey) generateMasterKey() error {
	_, err := io.ReadFull(rand.Reader, mkey.masterKey[:])
	if err != nil {
		return err
	}
	return nil
}

// derive the N-th id with current master key
func (mkey *MasterKey) deriveKey(id int) (dkey DerivedKey) {
	return
}
