package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/rivo/tview"
)

var keyGenWindowTitle = "- KEY GENERATION -"

func rawOutput(primitive tview.Primitive) *tview.Flex {
	return tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(nil, 0, 1, false).
			AddItem(primitive, 1, 1, true).
			AddItem(nil, 0, 1, false), 1, 1, true)
}

func showKeyGenPasswordPrompt(parent string, path string) {
	windowName := "showKeyGenPasswordPrompt"
	form := tview.NewForm()
	form.SetBorder(true)
	passwordField := tview.NewInputField().SetLabel("Password").
		SetFieldWidth(64).
		SetMaskCharacter('*')
	form.AddFormItem(passwordField)
	form.AddButton("OK", func() {
		err := masterKey.store([]byte(passwordField.GetText()), path)

		// display message after store
		if err != nil {
			showFailWindow("FAILURE", err.Error())
		} else {
			masterKey.path = path
			showSuccessWindow("SUCCESS", fmt.Sprint("Successfully Stored Master Key!!!\n", path), func() {
				info = infoWindow()
				mainFrame = mainFrameWindow()
				refreshBody()
				root.RemovePage(windowName)
				root.RemovePage(parent)
			})
		}
	})
	form.SetFocus(0)

	root.AddPage(windowName, popup(40, 10, form), true, true)
}

func showKeyGenWindow() {
	windowName := "showKeyGenWindow"
	text := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignLeft).
		SetWrap(false)

	// create a master key
	masterKey = newMasterKey()
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
		showKeyGenPasswordPrompt(windowName, inputField.GetText())
	})
	form.SetFocus(0)

	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow).
		SetTitle(keyGenWindowTitle).
		SetBorder(true)
	flex.AddItem(text, 0, 1, false)
	flex.AddItem(form, 0, 1, true)

	root.AddPage(windowName, popup(80, 15, flex), true, true)
}
