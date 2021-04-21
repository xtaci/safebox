package main

type IKeyExport interface {
	Name() string
	Export(key []byte) []byte
}
