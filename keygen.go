package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/rivo/tview"
)

func keyGenWindow() (content tview.Primitive) {
	text := tview.NewTextView().
		SetTextAlign(tview.AlignLeft).
		SetDynamicColors(true)

	// create a master key
	masterKey := newMasterKey()
	masterKey.generateMasterKey(nil)

	fmt.Fprint(text, "[red]Generate Master Key\n\n")
	fmt.Fprintf(text, "[blue::bl]%v...\n\n", hex.EncodeToString(masterKey.masterKey[:16]))
	fmt.Fprint(text, "[red::b]MAKE SURE YOU BACKUP THIS FILE CORRECTLY\n")
	fmt.Fprint(text, "[gray::]This file will be save to:")

	// path input field
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	form := tview.NewForm()
	form.AddInputField("Path", path, 32, nil, nil)
	form.AddButton("Save", func() {
	})
	form.AddButton("Cancel", nil)

	//
	grid := tview.NewGrid()
	grid.AddItem(text, 0, 0, 20, 20, 0, 0, false)
	grid.AddItem(form, 20, 0, 20, 20, 0, 0, false)
	return grid
}
