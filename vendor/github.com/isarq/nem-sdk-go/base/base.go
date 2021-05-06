package base

import (
	"fmt"
)

type Tx interface {
	GetTx() Transaction
}

type Transaction interface {
	GetType() int
	GetCommon() *CommonTransaction
	String() string
}

type Error struct {
	TimeStamp int64
	Error     string
	Message   string
	Status    int
}

// Before a mosaic can be created or transferred, a corresponding
// definition of the mosaic has to be created and published to the network.
// This is done via a mosaic definition creation transaction.
type MosaicDefinitionCreationTransaction struct {
	CommonTransaction
	CreationFee      float64          `json:"creationFee"`
	CreationFeeSink  string           `json:"creationFeeSink"`
	MosaicDefinition MosaicDefinition `json:"mosaicDefinition"`
}

// A mosaic definition describes an asset class. Some fields are mandatory
// while others are optional. The properties of a mosaic definition always
// have a default value and only need to be supplied if they differ from the default value.
type MosaicDefinition struct {
	Creator     string       `json:"creator"`
	Description string       `json:"description"`
	ID          MosaicID     `json:"id,omitempty"`
	Properties  []Properties `json:"properties"`
	Levy        Levy         `json:"levy,omitempty"`
}

type MosaicID struct {
	NamespaceID string `json:"namespaceId,omitempty"`
	Name        string `json:"name,omitempty"`
}

func (m MosaicID) String() string {
	return fmt.Sprintf(
		`{
			"NamespaceID": %v,
			"Name": %v }
		`,
		m.NamespaceID,
		m.Name)
}

// A mosaic describes an instance of a mosaic definition.
// Mosaics can be transferred by means of a transfer transaction.
type Mosaic struct {
	MosaicID MosaicID `json:"mosaicID,omitempty"`
	Quantity float64  `json:"quantity,omitempty"`
}

type Properties struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Levy struct {
	Type      int      `json:"type,omitempty"`
	Recipient string   `json:"recipient,omitempty"`
	MosaicID  MosaicID `json:"mosaicId,omitempty"`
	Fee       float64  `json:"fee,omitempty"`
}

type Node struct {
	Host string
	Port int
}

type MessageType struct {
	Value int
	Name  string
}

type ConsModif struct {
	ModificationType   int
	CosignatoryAccount string
}

type TransferTransaction struct {
	CommonTransaction
	Amount    float64  `json:"amount,omitempty"`
	Recipient string   `json:"recipient,omitempty"`
	Message   Message  `json:"message,omitempty"`
	Signature string   `json:"signature,omitempty"`
	Mosaics   []Mosaic `json:"mosaics,omitempty"`
}

type Common struct {
	Password, PrivateKey string
	IsHW                 bool
}

// Chain is the type of NEM chain.
type Chain struct {
	ID, Prefix int
	Char       string
}

type Data struct {
	Testnet, Mainnet, Mijin Chain
}

type Message struct {
	Type      int    `json:"type,omitempty"`
	Payload   string `json:"payload,omitempty"`
	PublicKey string `json:"publicKey,omitempty"`
}

type SignatureT struct {
	OtherHash struct {
		Data string
	}
	OtherAccount string
}

type Supply struct {
	Mosaic          string `json:"mosaic"`
	SupplyType      int    `json:"supplyType"`
	Delta           int    `json:"delta"`
	IsMultisig      bool   `json:"isMultisig"`
	MultisigAccount string `json:"multisigAccount"`
}

type MosaicsData struct {
	Quantity int      `json:"quantity"`
	MosaicID MosaicID `json:"mosaicId"`
}

func (m *MosaicsData) String() string {
	return fmt.Sprintf(
		`
			"Quantity": %v,
			"MosaicID": %v
		`,
		m.Quantity,
		m.MosaicID.String(),
	)
}

type MultisigAggregateModific struct {
	Modifications   []interface{} `json:"modifications"`
	RelativeChange  interface{}   `json:"relativeChange"`
	IsMultisig      bool          `json:"isMultisig"`
	MultisigAccount string        `json:"multisigAccount"`
}

type ImportanceTransfer struct {
	RemoteAccount   string `json:"remoteAccount"`
	Mode            int    `json:"mode"`
	IsMultisig      bool   `json:"isMultisig"`
	MultisigAccount string `json:"multisigAccount"`
}

type CommonTransaction struct {
	Type      int
	Version   int
	Signer    string
	TimeStamp *int64
	Fee       float64
	Deadline  *int64
}

func (c *CommonTransaction) String() string {
	return fmt.Sprintf(
		`
			"Type": %v,
			"Version": %v
			"Signer": %v
			"TimeStamp": %v
			"Fee": %v
			"Deadline": %v
		`,
		c.Type,
		c.Version,
		c.Signer,
		c.TimeStamp,
		c.Fee,
		c.Deadline)
}

type ProvisionNamespaceTransaction struct {
	CommonTransaction
	RentalFeeSink string  `json:"rentalFeeSink"`
	RentalFee     float64 `json:"rentalFee"`
	NewPart       string  `json:"newPart"`
	Parent        string  `json:"parent"`
}

type MultiSignSignatureTransaction struct {
	TimeStamp int64  `json:"timeStamp"`
	Fee       int    `json:"fee"`
	Type      int    `json:"type"`
	Deadline  int64  `json:"deadline"`
	Version   int    `json:"version"`
	Signer    string `json:"signer"`
	OtherHash struct {
		Data string `json:"data"`
	} `json:"otherHash"`
	OtherAccount string `json:"otherAccount"`
}

type TransactionResponse struct {
	TimeStamp  int64                           `json:"timeStamp"`
	Amount     float64                         `json:"amount"`
	Fee        float64                         `json:"fee"`
	Recipient  string                          `json:"recipient,omitempty"`
	Type       int                             `json:"type,omitempty"`
	Deadline   int64                           `json:"deadline"`
	Message    *Message                        `json:"message,omitempty"`
	Version    int                             `json:"version,omitempty"`
	Signer     string                          `json:"signer,omitempty"`
	OtherTrans Transaction                     `json:"otherTrans,omitempty"`
	Signatures []MultiSignSignatureTransaction `json:"signatures,omitempty"`
}

type MultiSignTransaction struct {
	CommonTransaction
	OtherTrans interface{}                     `json:"otherTrans"`
	Signatures []MultiSignSignatureTransaction `json:"signatures,omitempty"`
}

type AbstractTransaction struct {
	TimeStamp int64    `json:"timeStamp,omitempty"`
	Amount    float64  `json:"amount,omitempty"`
	Fee       float64  `json:"fee,omitempty"`
	Recipient string   `json:"recipient,omitempty"`
	Type      int      `json:"type,omitempty"`
	Deadline  int64    `json:"deadline,omitempty"`
	Message   *Message `json:"message,omitempty"`
	Version   int      `json:"version,omitempty"`
	Signer    string   `json:"signer,omitempty"`
}

type TransactionMosaic struct {
	AbstractTransaction
	Signature string        `json:"signature"`
	Mosaics   []MosaicsData `json:"mosaics"`
}

func (t *TransactionMosaic) GetType() int {
	return t.Type
}

func (t *TransactionMosaic) GetCommon() *CommonTransaction {
	return &CommonTransaction{
		Type:      t.Type,
		Version:   t.Version,
		Signer:    t.Signer,
		TimeStamp: &t.TimeStamp,
		Fee:       t.Fee,
		Deadline:  &t.Deadline,
	}
}

func (t *TransactionMosaic) String() string {
	return fmt.Sprintf(
		`
			"TimeStamp": %v,
			"Amount": %v,
			"Fee": %v,
			"Recipient": %v,
			"Type": %v,
			"Deadline": %v,
			"Message": %v,
			"Version": %v,
			"Signer": %v
		`,
		t.TimeStamp,
		t.Amount,
		t.Fee,
		t.Recipient,
		t.Type,
		t.Deadline,
		t.Message,
		t.Version,
		t.Signer,
	)
}

func (t *MosaicDefinitionCreationTransaction) GetType() int {
	mosaic := t.Type
	return mosaic
}

func (t *MosaicDefinitionCreationTransaction) String() string {
	return fmt.Sprintf(
		`
			"Common": %v,
			"CreationFee": %v,
			"CreationFeeSink": %v,
			"MosaicDefinition": %v
		`,
		t.CommonTransaction.String(),
		t.CreationFee,
		t.CreationFeeSink,
		t.MosaicDefinition,
	)
}

func (t *MosaicDefinitionCreationTransaction) GetCommon() *CommonTransaction {
	return &CommonTransaction{
		Type:      t.Type,
		Version:   t.Version,
		TimeStamp: t.TimeStamp,
		Deadline:  t.Deadline,
		Signer:    t.Signer,
		Fee:       t.Fee,
	}
}

func (t *MosaicDefinitionCreationTransaction) GetMosaic() MosaicDefinition {
	mosaic := t.MosaicDefinition
	return mosaic
}

func (t *MosaicDefinitionCreationTransaction) GetMosaicId() MosaicID {
	mosaic := t.MosaicDefinition.ID
	return mosaic
}

func (t *MosaicDefinitionCreationTransaction) GetMosaicTx() *MosaicDefinitionCreationTransaction {
	return t
}

func (t *MosaicDefinitionCreationTransaction) GetTx() Transaction {
	return t
}

func (t *ProvisionNamespaceTransaction) GetType() int {
	return t.Type
}

func (t *ProvisionNamespaceTransaction) GetCommon() *CommonTransaction {
	return &CommonTransaction{
		Type:      t.Type,
		Version:   t.Version,
		TimeStamp: t.TimeStamp,
		Deadline:  t.Deadline,
		Signer:    t.Signer,
		Fee:       t.Fee,
	}
}

func (t *ProvisionNamespaceTransaction) String() string {
	return fmt.Sprintf(
		`
			"Common": %v,
			"RentalFee": %v,
			"RentalFeeSink": %v,
			"Parent": %v,
			"Parent": %v
		`,
		t.CommonTransaction.String(),
		t.RentalFee,
		t.RentalFeeSink,
		t.Parent,
		t.NewPart,
	)
}

func (t *ProvisionNamespaceTransaction) GetTx() Transaction {
	return t
}

func (t *MultiSignTransaction) GetType() int {
	return t.Type
}

func (t *MultiSignTransaction) GetCommon() *CommonTransaction {
	return &CommonTransaction{
		Type:      t.Type,
		Version:   t.Version,
		TimeStamp: t.TimeStamp,
		Deadline:  t.Deadline,
		Signer:    t.Signer,
		Fee:       t.Fee,
	}
}

func (t *MultiSignTransaction) String() string {
	return fmt.Sprintf(
		`
			"Common": %v,
			"OtherTrans": %v,
			"Signatures": %v
		`,
		t.CommonTransaction.String(),
		t.OtherTrans,
		t.Signatures,
	)
}

func (t *TransferTransaction) GetType() int {
	return t.Type
}

func (t *TransferTransaction) GetCommon() *CommonTransaction {
	return &CommonTransaction{
		Type:      t.Type,
		Version:   t.Version,
		TimeStamp: t.TimeStamp,
		Deadline:  t.Deadline,
		Signer:    t.Signer,
		Fee:       t.Fee,
	}
}

func (t *TransferTransaction) String() string {
	return fmt.Sprintf(
		`
			"Common": %v,
			"Amount": %v,
			"Recipient": %v,
			"Signature": %v,
			"Message": %v
			"Mosaics": %v
		`,
		t.CommonTransaction.String(),
		t.Amount,
		t.Recipient,
		t.Signature,
		t.Message,
		t.Mosaics,
	)
}

func (t *TransferTransaction) GetTx() Transaction {
	return t
}
