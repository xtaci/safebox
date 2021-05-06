package model

import (
	"errors"
	"fmt"
	"github.com/isarq/nem-sdk-go/base"
)

// Supported predefined chains.
var Data = base.Data{
	Testnet: base.Chain{-104, 98, "T"},
	Mainnet: base.Chain{104, 68, "N"},
	Mijin:   base.Chain{96, 60, "M"},
}

// Gets a network prefix from network id
// param id - A network id
// return - The network prefix
func Id2Prefix(id int) byte {
	if id == 104 {
		return 0x68
	} else if id == -104 {
		return 0x98
	} else {
		return 0x60
	}
}

// Gets the starting char of the addresses of a network id
// param id - A network id
// return - The starting char of addresses
func Id2Char(id int) string {
	if id == 104 {
		return "N"
	} else if id == -104 {
		return "T"
	} else {
		return "M"
	}
}

// Gets the network id from the starting char of an address
// param startChar - A starting char from an address
// return - The network id
func Char2Id(startChar string) int {
	if startChar == "N" {
		return 104
	} else if startChar == "T" {
		return -104
	} else {
		return 96
	}
}

// Gets the network version
// param val - A version number (1 or 2)
// param network - A network id
// return A network version
func GetVersion(val, network int) int {
	if network == Data.Mainnet.ID {
		return 0x68000000 | val
	} else if network == Data.Testnet.ID {
		return 0x98000000 | val
	}
	return 0x60000000 | val
}

// NewChain parses byte value into a chain.
func NewChain(v int) (ch base.Chain, err error) {
	switch v {
	case -104:
		fmt.Println("Testnet: ", v)
		ch = Data.Testnet
	case 104:
		fmt.Println("Mainnet: ", v)
		ch = Data.Mainnet
	case 96:
		fmt.Println("Mijin: ", v)
		ch = Data.Mijin
	default:
		err = errors.New("core: invalid chain id")
	}
	return
}
