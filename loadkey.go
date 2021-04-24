package main

import (
	"fmt"
	"os"

	"github.com/rivo/tview"
)

func showLoadPassword(parent string, path string) {
	const (
		windowName   = "showLoadPassword"
		windowWidth  = 40
		windowHeight = 7
		windowTitle  = "- MASTERKEY DECRYPTION PASSWORD -"
	)

	form := tview.NewForm()
	form.SetTitle(windowTitle)
	form.SetBorder(true)
	passwordField := tview.NewInputField().SetLabel("Password").
		SetFieldWidth(64).
		SetMaskCharacter('*')

	form.AddButton("OK", func() {
		// create a master key
		masterKeyToLoad := newMasterKey()
		err := masterKeyToLoad.load([]byte(passwordField.GetText()), path)
		if err != nil {
			showFailWindow(err.Error())
			layoutRoot.RemovePage(windowName)
		} else {
			showSuccessWindow(fmt.Sprintf("Successfully Loaded Master Key!!!\n%v", path), func() {
				masterKey = masterKeyToLoad
				masterKey.path = path
				layoutInfo = infoWindow()
				layoutMainBody = mainFrameWindow()
				refreshBody()
				layoutRoot.RemovePage(windowName)
				layoutRoot.RemovePage(parent)
			})
		}
	})

	form.AddFormItem(passwordField)
	form.SetFocus(0)

	layoutRoot.AddPage(windowName, popup(windowWidth, windowHeight, form), true, true)
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
		SetText(path + "/.safebox.key").
		SetFieldWidth(64)

	form.AddFormItem(inputField)
	form.AddButton("Load", func() {
		showLoadPassword(windowName, inputField.GetText())
	})
	form.AddButton("...", func() {
		showDirWindow(inputField)
	})

	form.SetBorder(true)
	form.SetTitle(windowTitle)
	form.SetFocus(0)

	layoutRoot.AddPage(windowName, popup(windowWidth, windowHeight, form), true, true)
}
