package main

import (
	"fmt"
	"os"

	"github.com/rivo/tview"
)

var loadKeyWindowTitle = "-=- LOAD KEY -=-"

func passwordPromptLoad(path string) *tview.Flex {
	form := tview.NewForm()
	form.SetBorder(true)
	passwordField := tview.NewInputField().SetLabel("Password").
		SetFieldWidth(64).
		SetMaskCharacter('*')
	form.AddFormItem(passwordField)
	form.AddButton("OK", func() {
		// create a master key
		masterKeyToLoad := newMasterKey()
		err := masterKeyToLoad.load([]byte(passwordField.GetText()), path)
		if err != nil {
			root.AddAndSwitchToPage("prompt", failWindow(fmt.Sprintf("Failed Reading Master Key!!!\n%v", err)), true)
		} else {
			masterKey = masterKeyToLoad
			masterKey.path = path
			root.AddAndSwitchToPage("prompt", successWindow(fmt.Sprintf("Successfully Loaded Master Key!!!\n%v", path)), true)
		}
	})
	form.SetFocus(0)

	return modal(40, 10, form)
}

func loadKeyWindow() (content *tview.Flex) {
	text := tview.NewTextView().
		SetTextAlign(tview.AlignLeft).
		SetDynamicColors(true)

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
		root.AddAndSwitchToPage("prompt", passwordPromptLoad(inputField.GetText()), true)
	})
	form.AddButton("Cancel", nil)
	form.SetFocus(0)

	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow).
		SetBorder(true).
		SetTitle(loadKeyWindowTitle)
	flex.AddItem(text, 10, 1, false)
	flex.AddItem(form, 10, 1, true)

	return flex
}
