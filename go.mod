module github.com/xtaci/safebox

go 1.16

require (
	github.com/btcsuite/btcd v0.21.0-beta
	github.com/btcsuite/btcutil v1.0.2
	github.com/cosmos/cosmos-sdk v0.42.4
	github.com/cosmos/go-bip39 v1.0.0
	github.com/ethereum/go-ethereum v1.10.2
	github.com/gdamore/tcell/v2 v2.2.0
	github.com/rivo/tview v0.0.0-20210312174852-ae9464cc3598
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e
	github.com/spf13/cobra v1.1.3 // indirect
	github.com/stretchr/testify v1.7.0
	github.com/templexxx/xorsimd v0.4.1
	github.com/tendermint/tendermint v0.34.10 // indirect
	golang.org/x/crypto v0.0.0-20210415154028-4f45737414dc
)

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
