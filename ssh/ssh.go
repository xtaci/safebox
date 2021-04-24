package ssh

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha1"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"math/big"

	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/ssh"
)

type SSHExporter struct{}

func (exp *SSHExporter) Name() string {
	return "SSH(ECDSA)"
}
func (exp *SSHExporter) KeySize() int {
	return 32
}
func (exp *SSHExporter) Export(key []byte) ([]byte, error) {
	curve := elliptic.P256()
	// use pbkdf to extend the key
	if len(key) != curve.Params().BitSize/8 {
		keyLen := curve.Params().BitSize / 8
		key = pbkdf2.Key(key, []byte("SSH(EC)"), 1024, keyLen, sha1.New)
	}

	// Private Key generation
	var priv ecdsa.PrivateKey
	priv.Curve = curve // SSH uses this curve
	priv.D = new(big.Int).SetBytes(key)
	priv.PublicKey.X, priv.PublicKey.Y = priv.PublicKey.Curve.ScalarBaseMult(priv.D.Bytes())
	if priv.PublicKey.X == nil {
		return nil, errors.New("invalid private key")
	}

	bts, err := x509.MarshalECPrivateKey(&priv)
	if err != nil {
		return nil, err
	}

	// private
	privateKeyPEM := &pem.Block{Type: "EC PRIVATE KEY", Bytes: bts}
	var output bytes.Buffer
	if err := pem.Encode(&output, privateKeyPEM); err != nil {
		return nil, err
	}

	// public
	publicECKey, err := ssh.NewPublicKey(&priv.PublicKey)
	if err != nil {
		return nil, err
	}

	public := ssh.MarshalAuthorizedKey(publicECKey)

	output.Write([]byte("\n"))
	output.Write(public)

	return output.Bytes(), nil
}
