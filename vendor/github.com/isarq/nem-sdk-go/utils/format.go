package utils

import (
	"encoding/json"
	"fmt"
	"github.com/isarq/nem-sdk-go/base"
	"github.com/isarq/nem-sdk-go/extras"
)

// Convert a public key to NEM address
// param input - The account public key
// param networkId - The current network id
// return - A clean NEM address
func PubToAddress() {

}

func Struc2Json(data interface{}) string {
	var rawerr base.Error
	json.Unmarshal([]byte(fmt.Sprint(data)), &rawerr)
	if rawerr.Status == 0 {
		r, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			fmt.Println("Error:", err)
		}
		return string(r) + "\n_________________\n"
	}
	r, err := json.MarshalIndent(rawerr, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
	}
	return string(r) + "\n_________________\n"
}

// Return mosaic name from mosaicId object
// param mosaicId - A mosaicId object
// return The mosaic name
func MosaicIdToName(mosaicId base.MosaicID) string {
	if extras.IsEmpty(mosaicId) {
		return ""
	}
	return fmt.Sprintf("%v:%v", mosaicId.NamespaceID, mosaicId.Name)
}
