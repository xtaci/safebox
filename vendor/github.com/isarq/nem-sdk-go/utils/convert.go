package utils

import (
	"encoding/hex"
)

// Hex2BaReversed Reversed convertion of hex to []Bytes.
// param hexx - An hex string
// return []Bytes
func Hex2BaReversed(hexx string) []byte {
	data, err := hex.DecodeString(hexx)
	if err != nil {
		panic(err)
	}
	output := make([]byte, len(data))
	j := len(data) - 1
	for i := 0; i < len(data); i++ {
		output[j] = data[i]
		j--
	}
	return output
}

// Convert hex to []Bytes.
// param hexx - An hex string
// return bytes
func Hex2Bt(hexx string) []byte {
	data, err := hex.DecodeString(hexx)
	if err != nil {
		panic(err)
	}
	return data
}

// Convert an []Bytes to hex
// param b - An []Bytes
// return string
func Bt2Hex(b []byte) string {
	dst := make([]byte, hex.EncodedLen(len(b)))
	hex.Encode(dst, b)
	return string(dst)
}

// Convert hex to string
// param hexx - An hex string
// return string
func Hex2a(hexx string) string {
	bs, err := hex.DecodeString(hexx)
	if err != nil {
		panic(err)
	}
	return string(bs)
}

// Convert Bytes slice to UTF-8 string
// param ua - An slice Bytes
// return string
func BtToUtf8(ua []byte) string {
	return string(ua)
}

// Convert UTF-8 to hex
// @param str - An UTF-8 string
// @return string
func Utf8ToHex(str string) string {
	b := []byte(str)
	rawString := hex.EncodeToString(b)
	return rawString
}
