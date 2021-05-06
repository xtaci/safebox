package utils

import (
	"github.com/isarq/nem-sdk-go/base"
	"github.com/isarq/nem-sdk-go/extras"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Clean a text input amount and return it as number
// param n - The number as a string
// return The clean amount
func CleanTextAmount(n string) float64 {
	a := strings.Replace(n, ",", ".", 0)
	return extras.Number(a)
}

// Check if a text input amount is valid
// param n - The number as a string
// return True if valid, false otherwise
//func IsTextAmountValid(n string) bool {
//    // Force n as a string and replace decimal comma by a dot if any
//    var nn = Number(n.toString().replace(/,/g, '.'));
//    return !Number.isNaN(nn) && Number.isFinite(nn) && nn >= 0;
//}

// Convert an endpoint object to an endpoint url
// param endpoint - An endpoint object
// return An endpoint url
func FormatEndpoint(endpoint base.Node) string {
	port := strconv.Itoa(endpoint.Port)
	return endpoint.Host + ":" + port
}

// Check if a private key is valid
// param privatekey - A private key
// eturn True if valid, false otherwise
func IsPrivateKeyValid(privateKey string) bool {
	if len(privateKey) != 64 && len(privateKey) != 66 {
		log.Println("Private key length must be 64 or 66 characters !")
		return false
	} else if !IsHexadecimal(privateKey) {
		log.Println("Private key must be hexadecimal only !")
		return false
	} else {
		return true
	}
}

// Check if a public key is valid
// param publicKey - A public key
// return - True if valid, false otherwise
func IsPublicKeyValid(publicKey string) bool {
	if len(publicKey) != 64 {
		log.Println("Private key length must be 64 or 66 characters !")
		return false
	} else if !IsHexadecimal(publicKey) {
		log.Println("Private key must be hexadecimal only !")
		return false
	} else {
		return true
	}
}

// Test if a string is hexadecimal
// param str - A string to test
// return True if correct, false otherwise
func IsHexadecimal(str string) bool {
	exp := regexp.MustCompile("^[0-9a-fA-F]+$")
	if exp.MatchString(str) == true {
		return true
	}
	return false
}

// Create a time stamp for a NEM transaction
// NEM EPOCH = UTC(2015, 3, 29, 0, 6, 25, 0)
// return The NEM transaction time stamp in milliseconds
func CreateNEMTimeStamp() int64 {
	// 1427587585
	return int64(math.Floor(float64(time.Now().Unix() - 1427587585)))
}

// Fix a private key
// param privatekey - An hex private key
// return - The fixed hex private key
func FixPrivateKey(privateKey string) string {
	rest := "0000000000000000000000000000000000000000000000000000000000000000" + strings.Replace(privateKey, "^00",
		"", -1)
	return rest[len(rest)-64:]
}

// Mimics jQuery's grep function
func Grep(item []base.Properties) map[string]string {
	pro := make(map[string]string)
	for _, p := range item {
		pro[p.Name] = p.Value
	}
	return pro
}
