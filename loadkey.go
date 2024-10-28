package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/rivo/tview"
)

func showLoadPassword(parent string, path string) {
	const (
		windowName   = "showLoadPassword"
		windowWidth  = 40
		windowHeight = 6
		windowTitle  = "🔓MASTERKEY DECRYPTION"
	)

	form := tview.NewForm()
	form.SetTitle(windowTitle)
	form.SetBorder(true)
	passwordField := tview.NewInputField().SetLabel("Password").
		SetFieldWidth(0).
		SetMaskCharacter('*')

	passChangedFunc := func(text string) {
		// create a master key
		masterKeyToLoad := newMasterKey()
		err := masterKeyToLoad.load([]byte(passwordField.GetText()), path)
		if err == nil {
			showSuccessWindow(fmt.Sprintf("Successfully Decrypted Master Key!!!\n%v", path), func() {
				masterKey = masterKeyToLoad
				masterKey.path = path
				refresh()
				layoutRoot.RemovePage(windowName)
				layoutRoot.RemovePage(parent)
			})
		}
	}
	// compute hash on the fly
	passwordField.SetChangedFunc(passChangedFunc)

	form.AddFormItem(passwordField)
	form.SetFocus(0)

	layoutRoot.AddPage(windowName, popup(windowWidth, windowHeight, form), true, true)

	// test empty pass
	passChangedFunc("")
}

func showLoadKeyWindow() {
	const (
		windowName   = "showLoadKeyWindow"
		windowWidth  = 80
		windowHeight = 7
		windowTitle  = "-=- LOAD MASTER KEY -=-"
	)

	// path input field
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	form := tview.NewForm()
	// input field setting
	inputField := tview.NewInputField().
		SetLabel("Path: ").
		SetText(filepath.Join(path, ".safebox.key")).
		SetFieldWidth(0)

	form.AddFormItem(inputField)
	form.AddButton("Load", func() {
		if _, err := os.Stat(inputField.GetText()); os.IsNotExist(err) {
			showFailWindow(fmt.Sprintf("File: %v does not exists", inputField.GetText()))
		} else {
			showLoadPassword(windowName, inputField.GetText())
		}
	})
	form.AddButton("...", func() {
		showDirWindow(inputField)
	})

	form.SetBorder(true)
	form.SetTitle(windowTitle)
	form.SetFocus(0)

	layoutRoot.AddPage(windowName, popup(windowWidth, windowHeight, form), true, true)
}
