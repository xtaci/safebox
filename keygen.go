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

func showKeyGenPasswordPrompt(newkey *MasterKey, parent string, path string) {
	windowName := "showKeyGenPasswordPrompt"
	form := tview.NewForm()
	form.SetBorder(true)
	passwordField := tview.NewInputField().SetLabel("Password").
		SetFieldWidth(64).
		SetMaskCharacter('*')
	form.AddFormItem(passwordField)
	form.AddButton("OK", func() {
		newkey.changePassword([]byte(passwordField.GetText()))
		err := newkey.store(path)

		// display message after store
		if err != nil {
			showFailWindow("FAILURE", err.Error())
		} else {
			newkey.path = path
			showSuccessWindow("SUCCESS", fmt.Sprint("Successfully Stored Master Key!!!\n", path), func() {
				// set masterkey to newkey and update view
				masterKey = newkey
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
	newkey := newMasterKey()
	newkey.generateMasterKey(nil)

	fmt.Fprint(text, "GENERATED MASTER KEY:\n\n")
	fmt.Fprintf(text, "[darkorange::bl]%v...\n\n", hex.EncodeToString(newkey.masterKey[:32]))
	fmt.Fprint(text, "[darkorange::bu]MAKE SURE YOU BACKUP THIS FILE CORRECTLY\n")

	// path input field
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	form := tview.NewForm()
	inputField := tview.NewInputField().
		SetLabel("Save To:").
		SetText(path + "/.safebox.key").
		SetFieldWidth(64)
	form.AddFormItem(inputField)
	form.AddButton("Save", func() {
		// check file existence
		if _, err := os.Stat(inputField.GetText()); os.IsNotExist(err) {
			showKeyGenPasswordPrompt(newkey, windowName, inputField.GetText())
		} else {
			showFailWindow("FAILURE", "MASTER KEY FILE EXISTS, IF YOU WANT TO OVERWRITE, PLEASE DELETE THIS FILE BY YOURSELF.")
		}
	})
	form.AddButton("...", func() {
		showDirWindow(inputField)
	})

	form.SetFocus(0)

	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow).
		SetTitle(keyGenWindowTitle).
		SetBorder(true)
	flex.AddItem(text, 0, 1, false)
	flex.AddItem(form, 0, 1, true)

	root.AddPage(windowName, popup(100, 12, flex), true, true)
}
