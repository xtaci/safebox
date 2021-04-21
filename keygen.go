package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/rivo/tview"
)

var keyGenWindowTitle = "- KEY GENERATION -"

func keyGenWindow() (content *tview.Flex) {
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
	inputField := tview.NewInputField().
		SetLabel("Path: ").
		SetText(path + "/.safebox.key").
		SetFieldWidth(64)
	form.AddFormItem(inputField)
	form.AddButton("Save", func() {

	})
	form.AddButton("Cancel", nil)
	form.SetFocus(0)

	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow).
		SetBorder(true).
		SetTitle(keyGenWindowTitle)
	flex.AddItem(text, 0, 1, false)
	flex.AddItem(form, 0, 1, true)
	flex.AddItem(tview.NewBox(), 0, 8, false)

	return flex
}
