package main

import (
	"fmt"
	"os"

	"github.com/rivo/tview"
)

var loadKeyWindowTitle = "-=- LOAD MASTER KEY -=-"
var loadKeyPasswordTitle = "-=- PASSWORD TO DECRYPT MASTER KEY -=-"

func showLoadPassword(parent string, path string) {
	windowName := "showLoadPassword"
	form := tview.NewForm()
	form.SetBorder(true)
	form.SetTitle(loadKeyPasswordTitle)
	passwordField := tview.NewInputField().SetLabel("Password").
		SetFieldWidth(64).
		SetMaskCharacter('*')

	form.AddButton("OK", func() {
		// create a master key
		masterKeyToLoad := newMasterKey()
		err := masterKeyToLoad.load([]byte(passwordField.GetText()), path)
		if err != nil {
			showFailWindow("FAILURE", err.Error())
			layoutRoot.RemovePage(windowName)
		} else {
			showSuccessWindow("SUCCESS", fmt.Sprintf("Successfully Loaded Master Key!!!\n%v", path), func() {
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

	layoutRoot.AddPage(windowName, popup(40, 7, form), true, true)
}

func showLoadKeyWindow() {
	windowName := "showLoadKeyWindow"

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
	form.SetTitle(loadKeyWindowTitle)
	form.SetFocus(0)

	layoutRoot.AddPage(windowName, popup(80, 7, form), true, true)
}
