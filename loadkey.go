package main

import (
	"fmt"
	"os"

	"github.com/rivo/tview"
)

var loadKeyWindowTitle = "-=- LOAD MASTER KEY -=-"
var loadKeyPasswordTitle = "-=- PASSWORD TO DECRYPT MASTER KEY -=-"

func passwordPromptLoad(path string) *tview.Flex {
	form := tview.NewForm()
	form.SetBorder(true)
	form.SetTitle(loadKeyPasswordTitle)
	passwordField := tview.NewInputField().SetLabel("Password").
		SetFieldWidth(64).
		SetMaskCharacter('*')
	form.AddFormItem(passwordField)
	form.AddButton("OK", func() {
		// create a master key
		masterKeyToLoad := newMasterKey()
		err := masterKeyToLoad.load([]byte(passwordField.GetText()), path)
		if err != nil {
			addAndShowPopup("msgbox", failWindow(fmt.Sprintf("Failed Reading Master Key!!!\n%v", err)))
		} else {
			masterKey = masterKeyToLoad
			masterKey.path = path

			addAndShowPopup("msgbox", successWindow(fmt.Sprintf("Successfully Loaded Master Key!!!\n%v", path)))
		}
	})
	form.SetFocus(0)

	return modal(40, 10, form)
}

func loadKeyWindow() (content tview.Primitive) {
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
	form.AddButton("Load", func() {
		addAndShowPopup("load password promopt", passwordPromptLoad(inputField.GetText()))
	})
	form.AddButton("Cancel", nil)
	form.SetBorder(true)
	form.SetTitle(loadKeyWindowTitle)
	form.SetFocus(0)

	return form
}
