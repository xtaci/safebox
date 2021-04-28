package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
	"runtime"
	"time"
	"unsafe"

	xor "github.com/templexxx/xorsimd"
	"golang.org/x/crypto/pbkdf2"
)

var (
	ErrInvalidKeyId      = errors.New("invalid key id")
	ErrPasswordIncorrect = errors.New("incorrect password to decrypt the master key")
)

// common cryptograph settings in safebox
var (
	IV               = []byte{167, 79, 156, 18, 172, 27, 1, 164, 21, 242, 193, 252, 120, 230, 107, 115}
	SALT             = "safebox"
	pbkdf2Iterations = 4096
)

const (
	MasterKeyLength = 256 * 1024                      // 256 KB
	LabelSize       = 16                              // Maximum Label Size
	MaxKeys         = MasterKeyLength / aes.BlockSize // A const 16K
)

// record program startup time
var startTime time.Time

func init() {
	startTime = time.Now()
}

// MasterKey defines the master key memory structure
type MasterKey struct {
	password  []byte                // the password to encrypt storage masterkey
	createdAt int64                 // the date for masterkey creation
	path      string                // the path where the master key located
	masterKey [MasterKeyLength]byte // the mster key memory data(decrypted)
	labels    map[uint16]string     // labels of derived keys
}

// newMasterKey creates
func newMasterKey() *MasterKey {
	mkey := new(MasterKey)
	mkey.labels = make(map[uint16]string)
	return mkey
}

// generate master key with give entropy and system entropy
func (mkey *MasterKey) generateMasterKey(entropy []byte) error {
	_, err := io.ReadFull(rand.Reader, mkey.masterKey[:])
	if err != nil {
		return err
	}

	// entropy often comes from key strokes
	// overall entropys includes:
	//
	// 1. current unix time
	// 2. current process pid
	// 3. current os name
	// 4. given entropy
	// 5. current uid
	// 6. hostname
	// 7. workding dir
	// 8. process uptime

	var overallEntropy bytes.Buffer
	hostname, _ := os.Hostname()
	wd, _ := os.Getwd()
	fmt.Fprintf(&overallEntropy, "%v%v%v%v%v%v%v%v%v", time.Now().Unix(), time.Now().UnixNano(), os.Getpid(), runtime.GOOS, hostname, entropy, os.Getuid(), wd, int64(time.Since(startTime)))

	// we use all the entropy above to encrypt the master key again
	aesKey := sha256.Sum256(overallEntropy.Bytes())
	aesBlock, err := NewAESBlockCrypt(aesKey[:])
	if err != nil {
		return err
	}
	aesBlock.Encrypt(mkey.masterKey[:], mkey.masterKey[:])

	mkey.createdAt = time.Now().Unix()
	return nil
}

// derive the N-th id with current master key with specified key size
func (mkey *MasterKey) deriveKey(id uint16, keySize int) (key []byte, err error) {
	if id >= MaxKeys || id < 0 {
		return nil, ErrInvalidKeyId
	}

	// Approcach:
	// 1. take the N-th 16Byte as the key to encrypt the whole master key in CFB
	var encryptedKey [MasterKeyLength]byte
	aesKey := mkey.masterKey[int(id)*aes.BlockSize : (int(id)+1)*aes.BlockSize]
	aesBlock, err := NewAESBlockCrypt(aesKey)
	if err != nil {
		return nil, err
	}
	aesBlock.Encrypt(encryptedKey[:], mkey.masterKey[:])

	// 2. create a SHA-256 of the encrypted master key as the derived key
	md := sha256.Sum256(encryptedKey[:])

	// 3. use pbkdf2 to suit the key size
	if len(md) != keySize {
		key = pbkdf2.Key(md[:], []byte(SALT), pbkdf2Iterations, keySize, sha1.New)
	} else {
		key = md[:]
	}
	return key, err
}

// store the master key with password encryption
//
// Master key storage format:
//
// CREATED DATE (64bit unix timestamp)
// SHA256 of raw key(256 bit)
// Num Labels(16bit)
// ENCRYPTED MASTER KEY DATA(256KB)
// |Lable 1(uint16 little endian) | Label(16Byte 0 ending) |
// ...
// |Lable N(uint16 little endian) | Label(16Byte 0 ending) |
//
//
// NOTE: all integers are stored in LITTLE-ENDIAN
func (mkey *MasterKey) store(path string) (err error) {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Created At
	err = binary.Write(file, binary.LittleEndian, mkey.createdAt)
	if err != nil {
		return err
	}

	// SHA 256 of raw key
	md := sha256.Sum256(mkey.masterKey[:])
	_, err = file.Write(md[:])
	if err != nil {
		return err
	}

	// num labels
	err = binary.Write(file, binary.LittleEndian, uint16(len(mkey.labels)))
	if err != nil {
		return err
	}

	// write encrypted(AES-256) master key
	// expand the password to create AES-256 key
	key := pbkdf2.Key(mkey.password, []byte(SALT), pbkdf2Iterations, 32, sha1.New)
	var encryptedMasterKey [MasterKeyLength]byte
	aesBlock, err := NewAESBlockCrypt(key)
	if err != nil {
		return err
	}
	aesBlock.Encrypt(encryptedMasterKey[:], mkey.masterKey[:])

	// write encrypted key
	_, err = file.Write(encryptedMasterKey[:])
	if err != nil {
		return err
	}

	// write labels with 0 ending
	for id, label := range mkey.labels {
		var labelBytes [LabelSize]byte
		// write id
		err = binary.Write(file, binary.LittleEndian, id)
		if err != nil {
			return err
		}

		// write label(maximum 16 bytes)
		copy(labelBytes[:], []byte(label))
		_, err = file.Write(labelBytes[:])
		if err != nil {
			return err
		}
	}

	return nil
}

// load a master key from disk
func (mkey *MasterKey) load(password []byte, path string) (err error) {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	mkey.password = password
	// Read Created At
	err = binary.Read(file, binary.LittleEndian, &mkey.createdAt)
	if err != nil {
		return err
	}

	// Read SHA 256 of raw key
	var loadedMD [sha256.Size]byte
	_, err = io.ReadFull(file, loadedMD[:])
	if err != nil {
		return err
	}

	// read num labels
	var numLables uint16
	err = binary.Read(file, binary.LittleEndian, &numLables)
	if err != nil {
		return err
	}

	// Read encrypted Key
	_, err = io.ReadFull(file, mkey.masterKey[:])
	if err != nil {
		return err
	}

	// expand the password to create AES-256 key
	key := pbkdf2.Key(password, []byte(SALT), pbkdf2Iterations, 32, sha1.New)
	aesBlock, err := NewAESBlockCrypt(key)
	if err != nil {
		return err
	}

	// decrypt the key
	aesBlock.Decrypt(mkey.masterKey[:], mkey.masterKey[:])

	// Compare SHA256
	computedMD := sha256.Sum256(mkey.masterKey[:])
	if computedMD != loadedMD {
		return ErrPasswordIncorrect
	}

	// read lables
	for i := uint16(0); i < numLables; i++ {
		var lableBytes [LabelSize]byte
		var id uint16

		// read id
		err = binary.Read(file, binary.LittleEndian, &id)
		if err != nil {
			return err
		}

		// read label
		var label string
		_, err = io.ReadFull(file, lableBytes[:])
		if err != nil {
			return err
		}

		idx := bytes.IndexByte(lableBytes[:], 0)
		if idx == -1 {
			label = string(lableBytes[:])
		} else {
			label = string(lableBytes[:idx])
		}

		mkey.labels[id] = label
	}

	return nil
}

// set label
func (mkey *MasterKey) setLabel(idx uint16, label string) error {
	if idx >= MaxKeys || idx < 0 {
		return ErrInvalidKeyId
	}

	if label == "" {
		delete(mkey.labels, idx)
	} else {
		mkey.labels[idx] = label
	}
	return nil
}

// change password
func (mkey *MasterKey) changePassword(password []byte) {
	mkey.password = password
}

// BlockCrypt defines encryption/decryption methods for a given byte slice.
// Notes on implementing: the data to be encrypted contains a builtin
// nonce at the first 16 bytes
type BlockCrypt interface {
	// Encrypt encrypts the whole block in src into dst.
	// Dst and src may point at the same memory.
	Encrypt(dst, src []byte)

	// Decrypt decrypts the whole block in src into dst.
	// Dst and src may point at the same memory.
	Decrypt(dst, src []byte)
}

type aesBlockCrypt struct {
	encbuf [aes.BlockSize]byte
	decbuf [2 * aes.BlockSize]byte
	block  cipher.Block
}

// NewAESBlockCrypt https://en.wikipedia.org/wiki/Advanced_Encryption_Standard
func NewAESBlockCrypt(key []byte) (BlockCrypt, error) {
	c := new(aesBlockCrypt)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	c.block = block
	return c, nil
}

func (c *aesBlockCrypt) Encrypt(dst, src []byte) { encrypt(c.block, dst, src, c.encbuf[:]) }
func (c *aesBlockCrypt) Decrypt(dst, src []byte) { decrypt(c.block, dst, src, c.decbuf[:]) }

func encrypt(block cipher.Block, dst, src, buf []byte) {
	switch block.BlockSize() {
	case 8:
		encrypt8(block, dst, src, buf)
	case 16:
		encrypt16(block, dst, src, buf)
	default:
		panic("unsupported cipher block size")
	}
}

// optimized encryption for the ciphers which works in 8-bytes
func encrypt8(block cipher.Block, dst, src, buf []byte) {
	tbl := buf[:8]
	block.Encrypt(tbl, IV)
	n := len(src) / 8
	base := 0
	repeat := n / 8
	left := n % 8
	ptr_tbl := (*uint64)(unsafe.Pointer(&tbl[0]))

	for i := 0; i < repeat; i++ {
		s := src[base:][0:64]
		d := dst[base:][0:64]
		// 1
		*(*uint64)(unsafe.Pointer(&d[0])) = *(*uint64)(unsafe.Pointer(&s[0])) ^ *ptr_tbl
		block.Encrypt(tbl, d[0:8])
		// 2
		*(*uint64)(unsafe.Pointer(&d[8])) = *(*uint64)(unsafe.Pointer(&s[8])) ^ *ptr_tbl
		block.Encrypt(tbl, d[8:16])
		// 3
		*(*uint64)(unsafe.Pointer(&d[16])) = *(*uint64)(unsafe.Pointer(&s[16])) ^ *ptr_tbl
		block.Encrypt(tbl, d[16:24])
		// 4
		*(*uint64)(unsafe.Pointer(&d[24])) = *(*uint64)(unsafe.Pointer(&s[24])) ^ *ptr_tbl
		block.Encrypt(tbl, d[24:32])
		// 5
		*(*uint64)(unsafe.Pointer(&d[32])) = *(*uint64)(unsafe.Pointer(&s[32])) ^ *ptr_tbl
		block.Encrypt(tbl, d[32:40])
		// 6
		*(*uint64)(unsafe.Pointer(&d[40])) = *(*uint64)(unsafe.Pointer(&s[40])) ^ *ptr_tbl
		block.Encrypt(tbl, d[40:48])
		// 7
		*(*uint64)(unsafe.Pointer(&d[48])) = *(*uint64)(unsafe.Pointer(&s[48])) ^ *ptr_tbl
		block.Encrypt(tbl, d[48:56])
		// 8
		*(*uint64)(unsafe.Pointer(&d[56])) = *(*uint64)(unsafe.Pointer(&s[56])) ^ *ptr_tbl
		block.Encrypt(tbl, d[56:64])
		base += 64
	}

	switch left {
	case 7:
		*(*uint64)(unsafe.Pointer(&dst[base])) = *(*uint64)(unsafe.Pointer(&src[base])) ^ *ptr_tbl
		block.Encrypt(tbl, dst[base:])
		base += 8
		fallthrough
	case 6:
		*(*uint64)(unsafe.Pointer(&dst[base])) = *(*uint64)(unsafe.Pointer(&src[base])) ^ *ptr_tbl
		block.Encrypt(tbl, dst[base:])
		base += 8
		fallthrough
	case 5:
		*(*uint64)(unsafe.Pointer(&dst[base])) = *(*uint64)(unsafe.Pointer(&src[base])) ^ *ptr_tbl
		block.Encrypt(tbl, dst[base:])
		base += 8
		fallthrough
	case 4:
		*(*uint64)(unsafe.Pointer(&dst[base])) = *(*uint64)(unsafe.Pointer(&src[base])) ^ *ptr_tbl
		block.Encrypt(tbl, dst[base:])
		base += 8
		fallthrough
	case 3:
		*(*uint64)(unsafe.Pointer(&dst[base])) = *(*uint64)(unsafe.Pointer(&src[base])) ^ *ptr_tbl
		block.Encrypt(tbl, dst[base:])
		base += 8
		fallthrough
	case 2:
		*(*uint64)(unsafe.Pointer(&dst[base])) = *(*uint64)(unsafe.Pointer(&src[base])) ^ *ptr_tbl
		block.Encrypt(tbl, dst[base:])
		base += 8
		fallthrough
	case 1:
		*(*uint64)(unsafe.Pointer(&dst[base])) = *(*uint64)(unsafe.Pointer(&src[base])) ^ *ptr_tbl
		block.Encrypt(tbl, dst[base:])
		base += 8
		fallthrough
	case 0:
		xorBytes(dst[base:], src[base:], tbl)
	}
}

// optimized encryption for the ciphers which works in 16-bytes
func encrypt16(block cipher.Block, dst, src, buf []byte) {
	tbl := buf[:16]
	block.Encrypt(tbl, IV)
	n := len(src) / 16
	base := 0
	repeat := n / 8
	left := n % 8
	for i := 0; i < repeat; i++ {
		s := src[base:][0:128]
		d := dst[base:][0:128]
		// 1
		xor.Bytes16Align(d[0:16], s[0:16], tbl)
		block.Encrypt(tbl, d[0:16])
		// 2
		xor.Bytes16Align(d[16:32], s[16:32], tbl)
		block.Encrypt(tbl, d[16:32])
		// 3
		xor.Bytes16Align(d[32:48], s[32:48], tbl)
		block.Encrypt(tbl, d[32:48])
		// 4
		xor.Bytes16Align(d[48:64], s[48:64], tbl)
		block.Encrypt(tbl, d[48:64])
		// 5
		xor.Bytes16Align(d[64:80], s[64:80], tbl)
		block.Encrypt(tbl, d[64:80])
		// 6
		xor.Bytes16Align(d[80:96], s[80:96], tbl)
		block.Encrypt(tbl, d[80:96])
		// 7
		xor.Bytes16Align(d[96:112], s[96:112], tbl)
		block.Encrypt(tbl, d[96:112])
		// 8
		xor.Bytes16Align(d[112:128], s[112:128], tbl)
		block.Encrypt(tbl, d[112:128])
		base += 128
	}

	switch left {
	case 7:
		xor.Bytes16Align(dst[base:], src[base:], tbl)
		block.Encrypt(tbl, dst[base:])
		base += 16
		fallthrough
	case 6:
		xor.Bytes16Align(dst[base:], src[base:], tbl)
		block.Encrypt(tbl, dst[base:])
		base += 16
		fallthrough
	case 5:
		xor.Bytes16Align(dst[base:], src[base:], tbl)
		block.Encrypt(tbl, dst[base:])
		base += 16
		fallthrough
	case 4:
		xor.Bytes16Align(dst[base:], src[base:], tbl)
		block.Encrypt(tbl, dst[base:])
		base += 16
		fallthrough
	case 3:
		xor.Bytes16Align(dst[base:], src[base:], tbl)
		block.Encrypt(tbl, dst[base:])
		base += 16
		fallthrough
	case 2:
		xor.Bytes16Align(dst[base:], src[base:], tbl)
		block.Encrypt(tbl, dst[base:])
		base += 16
		fallthrough
	case 1:
		xor.Bytes16Align(dst[base:], src[base:], tbl)
		block.Encrypt(tbl, dst[base:])
		base += 16
		fallthrough
	case 0:
		xorBytes(dst[base:], src[base:], tbl)
	}
}

// decryption
func decrypt(block cipher.Block, dst, src, buf []byte) {
	switch block.BlockSize() {
	case 8:
		decrypt8(block, dst, src, buf)
	case 16:
		decrypt16(block, dst, src, buf)
	default:
		panic("unsupported cipher block size")
	}
}

// decrypt 8 bytes block, all byte slices are supposed to be 64bit aligned
func decrypt8(block cipher.Block, dst, src, buf []byte) {
	tbl := buf[0:8]
	next := buf[8:16]
	block.Encrypt(tbl, IV)
	n := len(src) / 8
	base := 0
	repeat := n / 8
	left := n % 8
	ptr_tbl := (*uint64)(unsafe.Pointer(&tbl[0]))
	ptr_next := (*uint64)(unsafe.Pointer(&next[0]))

	for i := 0; i < repeat; i++ {
		s := src[base:][0:64]
		d := dst[base:][0:64]
		// 1
		block.Encrypt(next, s[0:8])
		*(*uint64)(unsafe.Pointer(&d[0])) = *(*uint64)(unsafe.Pointer(&s[0])) ^ *ptr_tbl
		// 2
		block.Encrypt(tbl, s[8:16])
		*(*uint64)(unsafe.Pointer(&d[8])) = *(*uint64)(unsafe.Pointer(&s[8])) ^ *ptr_next
		// 3
		block.Encrypt(next, s[16:24])
		*(*uint64)(unsafe.Pointer(&d[16])) = *(*uint64)(unsafe.Pointer(&s[16])) ^ *ptr_tbl
		// 4
		block.Encrypt(tbl, s[24:32])
		*(*uint64)(unsafe.Pointer(&d[24])) = *(*uint64)(unsafe.Pointer(&s[24])) ^ *ptr_next
		// 5
		block.Encrypt(next, s[32:40])
		*(*uint64)(unsafe.Pointer(&d[32])) = *(*uint64)(unsafe.Pointer(&s[32])) ^ *ptr_tbl
		// 6
		block.Encrypt(tbl, s[40:48])
		*(*uint64)(unsafe.Pointer(&d[40])) = *(*uint64)(unsafe.Pointer(&s[40])) ^ *ptr_next
		// 7
		block.Encrypt(next, s[48:56])
		*(*uint64)(unsafe.Pointer(&d[48])) = *(*uint64)(unsafe.Pointer(&s[48])) ^ *ptr_tbl
		// 8
		block.Encrypt(tbl, s[56:64])
		*(*uint64)(unsafe.Pointer(&d[56])) = *(*uint64)(unsafe.Pointer(&s[56])) ^ *ptr_next
		base += 64
	}

	switch left {
	case 7:
		block.Encrypt(next, src[base:])
		*(*uint64)(unsafe.Pointer(&dst[base])) = *(*uint64)(unsafe.Pointer(&src[base])) ^ *(*uint64)(unsafe.Pointer(&tbl[0]))
		tbl, next = next, tbl
		base += 8
		fallthrough
	case 6:
		block.Encrypt(next, src[base:])
		*(*uint64)(unsafe.Pointer(&dst[base])) = *(*uint64)(unsafe.Pointer(&src[base])) ^ *(*uint64)(unsafe.Pointer(&tbl[0]))
		tbl, next = next, tbl
		base += 8
		fallthrough
	case 5:
		block.Encrypt(next, src[base:])
		*(*uint64)(unsafe.Pointer(&dst[base])) = *(*uint64)(unsafe.Pointer(&src[base])) ^ *(*uint64)(unsafe.Pointer(&tbl[0]))
		tbl, next = next, tbl
		base += 8
		fallthrough
	case 4:
		block.Encrypt(next, src[base:])
		*(*uint64)(unsafe.Pointer(&dst[base])) = *(*uint64)(unsafe.Pointer(&src[base])) ^ *(*uint64)(unsafe.Pointer(&tbl[0]))
		tbl, next = next, tbl
		base += 8
		fallthrough
	case 3:
		block.Encrypt(next, src[base:])
		*(*uint64)(unsafe.Pointer(&dst[base])) = *(*uint64)(unsafe.Pointer(&src[base])) ^ *(*uint64)(unsafe.Pointer(&tbl[0]))
		tbl, next = next, tbl
		base += 8
		fallthrough
	case 2:
		block.Encrypt(next, src[base:])
		*(*uint64)(unsafe.Pointer(&dst[base])) = *(*uint64)(unsafe.Pointer(&src[base])) ^ *(*uint64)(unsafe.Pointer(&tbl[0]))
		tbl, next = next, tbl
		base += 8
		fallthrough
	case 1:
		block.Encrypt(next, src[base:])
		*(*uint64)(unsafe.Pointer(&dst[base])) = *(*uint64)(unsafe.Pointer(&src[base])) ^ *(*uint64)(unsafe.Pointer(&tbl[0]))
		tbl, next = next, tbl
		base += 8
		fallthrough
	case 0:
		xorBytes(dst[base:], src[base:], tbl)
	}
}

func decrypt16(block cipher.Block, dst, src, buf []byte) {
	tbl := buf[0:16]
	next := buf[16:32]
	block.Encrypt(tbl, IV)
	n := len(src) / 16
	base := 0
	repeat := n / 8
	left := n % 8
	for i := 0; i < repeat; i++ {
		s := src[base:][0:128]
		d := dst[base:][0:128]
		// 1
		block.Encrypt(next, s[0:16])
		xor.Bytes16Align(d[0:16], s[0:16], tbl)
		// 2
		block.Encrypt(tbl, s[16:32])
		xor.Bytes16Align(d[16:32], s[16:32], next)
		// 3
		block.Encrypt(next, s[32:48])
		xor.Bytes16Align(d[32:48], s[32:48], tbl)
		// 4
		block.Encrypt(tbl, s[48:64])
		xor.Bytes16Align(d[48:64], s[48:64], next)
		// 5
		block.Encrypt(next, s[64:80])
		xor.Bytes16Align(d[64:80], s[64:80], tbl)
		// 6
		block.Encrypt(tbl, s[80:96])
		xor.Bytes16Align(d[80:96], s[80:96], next)
		// 7
		block.Encrypt(next, s[96:112])
		xor.Bytes16Align(d[96:112], s[96:112], tbl)
		// 8
		block.Encrypt(tbl, s[112:128])
		xor.Bytes16Align(d[112:128], s[112:128], next)
		base += 128
	}

	switch left {
	case 7:
		block.Encrypt(next, src[base:])
		xor.Bytes16Align(dst[base:], src[base:], tbl)
		tbl, next = next, tbl
		base += 16
		fallthrough
	case 6:
		block.Encrypt(next, src[base:])
		xor.Bytes16Align(dst[base:], src[base:], tbl)
		tbl, next = next, tbl
		base += 16
		fallthrough
	case 5:
		block.Encrypt(next, src[base:])
		xor.Bytes16Align(dst[base:], src[base:], tbl)
		tbl, next = next, tbl
		base += 16
		fallthrough
	case 4:
		block.Encrypt(next, src[base:])
		xor.Bytes16Align(dst[base:], src[base:], tbl)
		tbl, next = next, tbl
		base += 16
		fallthrough
	case 3:
		block.Encrypt(next, src[base:])
		xor.Bytes16Align(dst[base:], src[base:], tbl)
		tbl, next = next, tbl
		base += 16
		fallthrough
	case 2:
		block.Encrypt(next, src[base:])
		xor.Bytes16Align(dst[base:], src[base:], tbl)
		tbl, next = next, tbl
		base += 16
		fallthrough
	case 1:
		block.Encrypt(next, src[base:])
		xor.Bytes16Align(dst[base:], src[base:], tbl)
		tbl, next = next, tbl
		base += 16
		fallthrough
	case 0:
		xorBytes(dst[base:], src[base:], tbl)
	}
}

// per bytes xors
func xorBytes(dst, a, b []byte) int {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	if n == 0 {
		return 0
	}

	for i := 0; i < n; i++ {
		dst[i] = a[i] ^ b[i]
	}
	return n
}
