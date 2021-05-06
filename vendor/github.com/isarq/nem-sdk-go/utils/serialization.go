package utils

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/isarq/nem-sdk-go/base"
	"github.com/isarq/nem-sdk-go/extras"
	"math"
	"sort"
)

const (
	Const4bytessigner  = 32
	Const4bytesaddress = 40
)

type Temporary struct {
	entity base.Mosaic
	value  string
}

// The mosaic definition transaction type
const Mosaicdefinition = 0x4001 // 16385

// The provision namespace transaction type
const ProvisionNamespace = 0x2001 // 8193

// The multisignature signature transaction type
const MultisigSignature = 0x1002 // 4098

// safe string - each char is 8 bit */
func serializeSafeString(str string) []byte {
	var data []byte
	if str == "" {
		return encodeByte4(0xffffffff)
	}
	data = append(data, encodeByte4(len(str))...)
	j := hex.EncodeToString([]byte(str))
	e, _ := hex.DecodeString(j)
	data = append(data, e...)
	return data
}

func serializeUaString(str []byte) []byte {
	var data []byte
	if str == nil {
		return encodeByte4(0xffffffff)
	}
	data = append(data, encodeByte4(len(str))...)
	j := hex.EncodeToString(str)
	e, _ := hex.DecodeString(j)
	data = append(data, e...)
	return data
}

func serializeLong(value float64) []byte {
	var data []byte
	data = append(data, encodeByte4(value)...)
	data = append(data, encodeByte4(math.Floor(value/0x100000000))...)
	return data
}

// Mosaic id structure
func serializeMosaicId(entity base.MosaicID) []byte {
	var data []byte
	serializedNamespaceId := serializeSafeString(entity.NamespaceID)
	serializedName := serializeSafeString(entity.Name)

	data = append(data, encodeByte4(len(serializedNamespaceId)+len(serializedName))...)

	data = append(data, serializedNamespaceId...)

	data = append(data, serializedName...)
	return data
}

func serializeMosaicAndQuantity(mosaic base.Mosaic) []byte {
	var data []byte

	serializedMosaicId := serializeMosaicId(mosaic.MosaicID)
	serializedQuantity := serializeLong(mosaic.Quantity)

	data = append(data, encodeByte4(len(serializedMosaicId)+len(serializedQuantity))...)
	data = append(data, serializedMosaicId...)
	data = append(data, serializedQuantity...)

	return data
}

func serializeMosaics(entity []base.Mosaic) []byte {
	var data []byte
	var temp []Temporary

	data = append(data, encodeByte4(len(entity))...)

	var t Temporary
	for _, b := range entity {
		t.entity = b
		t.value = MosaicIdToName(b.MosaicID) + " : " + fmt.Sprint(b.Quantity)
		temp = append(temp, t)
	}

	sort.SliceStable(temp, func(i, j int) bool { return temp[i].value < temp[j].value })

	for i := 0; i < len(temp); i++ {
		entity := temp[i].entity
		serializedMosaic := serializeMosaicAndQuantity(entity)
		data = append(data, serializedMosaic...)
	}
	return data
}

func serializeProperty(entity base.Properties) []byte {
	var data []byte
	serializedName := serializeSafeString(entity.Name)
	serializedValue := serializeSafeString(entity.Value)

	data = append(data, encodeByte4(len(serializedName)+len(serializedValue))...)

	data = append(data, serializedName...)
	data = append(data, serializedValue...)

	return data
}

func serializeProperties(entity []base.Properties) []byte {
	var data []byte

	data = append(data, encodeByte4(len(entity))...)
	sort.SliceStable(entity, func(i, j int) bool { return entity[i].Name < entity[j].Name })

	for _, b := range entity {
		var entity base.Properties
		entity.Name = b.Name
		entity.Value = b.Value

		serializedProperty := serializeProperty(entity)
		data = append(data, serializedProperty...)
		//fmt.Printf("CANTIDAD: %v \n", data)
	}
	return data
}

func serializeLevy(entity base.Levy) []byte {
	var data []byte
	if extras.IsEmpty(entity) {
		data = append(data, encodeByte4(0)...)
		return data
	}

	// Length of recipient address field (always 40): 4 bytes (integer).
	// Example: 0x28, 0x00, 0x00, 0x00
	temp := serializeSafeString(entity.Recipient)

	// Length of mosaic id structure: 4 bytes (integer).
	// Example: 0x10, 0x00, 0x00, 0x00
	serializedMosaicId := serializeMosaicId(entity.MosaicID)

	serializedFee := serializeLong(entity.Fee)

	// Length of levy structure: 4 bytes (integer).
	// Example: 0x4c, 0x00, 0x00, 0x00
	data = append(data, encodeByte4(4+len(temp)+len(serializedMosaicId)+8)...)

	// Fee type: 4 bytes (integer). The following fee types are supported.
	// 0x01 (absolute fee)
	// 0x02 (percentile fee)
	// Example: 0x01, 0x00, 0x00, 0x00
	data = append(data, encodeByte4(entity.Type)...)

	data = append(data, temp...)

	data = append(data, serializedMosaicId...)

	data = append(data, serializedFee...)

	return data
}

// Mosaic definition creation transaction part
func serializeMosaicDefinition(entity base.MosaicDefinition) []byte {
	var data []byte
	Creator, _ := hex.DecodeString(entity.Creator)
	// Length of creator's public key byte array (always 32): 4 bytes (integer).
	// Always: 0x20, 0x00, 0x00, 0x00
	data = append(data, encodeByte4(len(Creator))...)
	//fmt.Printf("Len_Creator: 0x%v \n", hex.EncodeToString(encodeByte4(len(Creator))))
	// Public key bytes of creator: 32 bytes.
	data = append(data, Creator...)
	//fmt.Println("Creator: ", Creator)
	//fmt.Println("Creator: ", len(Creator))
	//fmt.Printf("Creator: 0x%v \n", hex.EncodeToString([]byte(Creator)))
	serializedMosaicId := serializeMosaicId(entity.ID)
	//fmt.Println("serializedMosaicId: ", len(serializedMosaicId))

	// Length of mosaic id structure: 4 bytes (integer).
	// Example: 0x0e, 0x00, 0x00, 0x00
	data = append(data, serializedMosaicId...)
	//fmt.Printf("Len_Mosaic_Id_Struct: 0x%v \n", hex.EncodeToString(encodeByte4(len(serializedMosaicId))))
	//
	//data = append(data, encodeByte4(serializedMosaicId)...)
	//fmt.Println("DATA 04: ", data)
	// Length of mosaic name string
	// Example: 0x15, 0x00, 0x00, 0x00

	// Mosaic name string: UTF8 encoded string.
	utf8ToUa := Hex2Bt(Utf8ToHex(entity.Description))
	temp := serializeUaString(utf8ToUa)
	//fmt.Println("utf8ToUa: ", utf8ToUa)
	data = append(data, temp...)

	temp = serializeProperties(entity.Properties)
	data = append(data, temp...)

	levy := serializeLevy(entity.Levy)
	data = append(data, levy...)

	return data
}

// Serialize a transaction struct
// param entity - A transaction struct
// return The serialized transaction
func SerializeTransaction(entity interface{}) []byte {
	var data []byte

	switch entity.(type) {
	case *base.TransferTransaction:
		tx, _ := entity.(*base.TransferTransaction)
		common, _ := commonHeader(tx)
		data = common

		data = append(data, encodeByte4(Const4bytesaddress)...) //const 28000000

		data = append(data, []byte(tx.Recipient)...) //signer 40 bytes

		data = append(data, encodeByte4(tx.Amount)...)                         //amount 0x40420f0000000000
		data = append(data, encodeByte4(math.Floor(tx.Amount/0x100000000))...) //amount 0x40420f0000000000

		if !extras.IsEmpty(tx.Message.Payload) && len(tx.Message.Payload) > 0 {
			msglength := len(tx.Message.Payload) / 2
			m, _ := hex.DecodeString(tx.Message.Payload)
			msg := string(m)

			data = append(data, encodeByte4(msglength+8)...) //msg body 0x23000000

			data = append(data, encodeByte4(tx.Message.Type)...) //msg type 0x01000000

			data = append(data, encodeByte4(msglength)...) // msg length 0x1b000000

			data = append(data, msg...) // msg payload 27
		} else {
			data = append(data, encodeByte4(0)...)
		}

		entityVersion := tx.Version & 0xffffff

		if entityVersion >= 2 {
			temp := serializeMosaics(tx.Mosaics)
			data = append(data, temp...)
		}

	// Mosaic Definition Creation transaction
	case *base.MosaicDefinitionCreationTransaction:
		//fmt.Println("Mosaicdefinition")

		tx, _ := entity.(*base.MosaicDefinitionCreationTransaction)
		common, _ := commonHeader(tx)
		data = common
		s := tx.GetMosaicTx()
		temp := serializeMosaicDefinition(s.GetMosaic())

		data = append(data, encodeByte4(len(temp))...)

		data = append(data, temp...)

		temp = serializeSafeString(s.CreationFeeSink)
		data = append(data, temp...)

		temp = serializeLong(s.CreationFee)
		data = append(data, temp...)

		// Provision Namespace transaction
	case *base.ProvisionNamespaceTransaction:
		//fmt.Println("ProvisionNamespace")
		tx, _ := entity.(*base.ProvisionNamespaceTransaction)

		common, _ := commonHeader(tx)
		data = common

		data = append(data, encodeByte4(len(tx.RentalFeeSink))...)
		// TODO: check that len(entity.RentalFee) is always 40 bytes
		data = append(data, tx.RentalFeeSink...)

		data = append(data, encodeByte4(tx.RentalFee)...)
		data = append(data, encodeByte4(math.Floor(tx.RentalFee/0x100000000))...)

		temp := serializeSafeString(tx.NewPart)
		data = append(data, temp...)

		temp = serializeSafeString(tx.Parent)
		data = append(data, temp...)

		// MultiSign wrapped transaction
	case *base.MultiSignTransaction:
		//fmt.Println("MultiSignSignature")
		tx, _ := entity.(*base.MultiSignTransaction)
		common, _ := commonHeader(tx)
		data = common
		temp := SerializeTransaction(tx.OtherTrans)

		data = append(data, encodeByte4(len(temp))...)
		data = append(data, temp...)
	}
	return data
}

func commonHeader(txstruct base.Transaction) (ch []byte, err error) {
	var data []byte
	tx := txstruct.GetCommon()
	Signer, _ := hex.DecodeString(tx.Signer)
	data = encodeByte4(tx.Type) //tx type 0x0101

	data = append(data, encodeByte4(tx.Version)...)    //version 0x01000098
	data = append(data, encodeByte4(*tx.TimeStamp)...) //timestamp 0xccaa7704

	data = append(data, encodeByte4(Const4bytessigner)...) //const 0x20000000

	data = append(data, Signer...)                    //signer 32 bytes
	data = append(data, encodeByte8(tx.Fee)...)       //fee 0xa086010000000000
	data = append(data, encodeByte4(*tx.Deadline)...) //deadline 0xdcb87704
	return data, nil
}

func encodeByte4(valor interface{}) []byte {
	var Type = make([]byte, 4)

	if s, ok := valor.(int); ok {
		binary.LittleEndian.PutUint32(Type, uint32(s))
	}
	if s, ok := valor.(int64); ok {
		binary.LittleEndian.PutUint32(Type, uint32(s))
	}
	if s, ok := valor.(float64); ok {
		binary.LittleEndian.PutUint32(Type, uint32(s))
	}
	return Type
}

func encodeByte8(valor interface{}) []byte {
	var Type = make([]byte, 8)

	if s, ok := valor.(int); ok {
		binary.LittleEndian.PutUint32(Type, uint32(s))
	}
	if s, ok := valor.(int64); ok {
		binary.LittleEndian.PutUint32(Type, uint32(s))
	}
	if s, ok := valor.(float64); ok {
		binary.LittleEndian.PutUint32(Type, uint32(s))
	}
	return Type
}
