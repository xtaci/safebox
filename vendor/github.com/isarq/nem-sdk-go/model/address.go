package model

import (
	"bytes"
	"crypto/rand"
	"encoding/base32"
	"encoding/hex"
	"fmt"
	"github.com/isarq/nem-sdk-go/external/crypto/ed25519"
	"github.com/isarq/nem-sdk-go/external/crypto/sha3"
	"github.com/isarq/nem-sdk-go/utils"
	"golang.org/x/crypto/ripemd160"
	"io"
	"strings"
)

const (
	// PrivateBytes stores the private key length in bytes.
	PrivateBytes = 32
	// PublicBytes stores the public key length in bytes.
	PublicBytes = 32
)

// KeyPair is a private/public crypto key pair.
type KeyPair struct {
	Private []byte
	Public  []byte
}

func (k *KeyPair) PublicString() string {
	return utils.Bt2Hex(k.Public)
}

func (k *KeyPair) PrivateString() string {
	return utils.Bt2Hex(k.Private)
}

// Sign signs the message with privateKey and returns a signature. It will
// panic if len(privateKey) is not PrivateKeySize.
func (k *KeyPair) Sign(msg string) []byte {
	_, pr, _ := ed25519.GenerateKey(bytes.NewReader(k.Private))
	return ed25519.Sign(pr, []byte(msg))
}

// Verify reports whether sig is a valid signature of message by publicKey. It
// will panic if len(publicKey) is not PublicKeySize.
func Verify(publicKey, message, sig []byte) bool {
	if ed25519.Verify(publicKey, message, sig) {
		return true
	}
	return false
}

// ToAddress convert a public key to a NEM address
// param publicKey - A public key
// param networkId - A network id
// return - The NEM address
func ToAddress(publicKey string, networkId int) string {
	pk, err := hex.DecodeString(strings.TrimSpace(publicKey)) //Python
	if err != nil {
		panic(err)
	}

	h := sha3.SumKeccak256(pk)
	networkPrefix := Id2Prefix(networkId)

	md := ripemd160.New()
	md.Write(h[:])

	s := md.Sum(nil)

	s = append([]byte{networkPrefix}, s...)
	h = sha3.SumKeccak256(s)

	address := append(s, h[:4]...)
	return base32.StdEncoding.EncodeToString(address)
}

// KeyPairCreate generates a KeyPair using specified string 32 length or empty
// param pk - A string 32 length or nil
// return - The NEM KeyPair
func KeyPairCreate(pk string) KeyPair {
	if pk != "" {
		if len(pk) != 64 {
			err := fmt.Errorf("insufficient seed length, should be %d, but got %d", 64, len(pk))
			panic(err)
		}
		seed, err := hex.DecodeString(strings.TrimSpace(pk))
		if err != nil {
			panic(err)
		}
		pair, err := FromSeed(seed)
		if err != nil {
			panic(err)
		}
		return pair
	}

	seed := make([]byte, PrivateBytes)
	_, err := io.ReadFull(rand.Reader, seed)
	if err != nil {
		panic(err)
	}
	pair, err := FromSeed(seed)
	if err != nil {
		panic(err)
	}
	return pair
}

// FromSeed generates a new private/public key pair using specified private key.
// param pk - A PrivateBytes
// return - The NEM KeyPair
func FromSeed(seed []byte) (KeyPair, error) {
	if len(seed) != PrivateBytes {
		return KeyPair{},
			fmt.Errorf("insufficient seed length, should be %d, but got %d", 64, len(seed))
	}
	pub, pr, err := ed25519.GenerateKey(bytes.NewReader(seed))
	if err != nil {
		return KeyPair{}, err
	}
	return KeyPair{pr[:PrivateBytes], pub}, nil
}
