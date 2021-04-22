package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var keyGenWindowTitle = "- KEY GENERATION -"

func modal(width int, height int, primitive tview.Primitive) *tview.Flex {
	return tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(nil, 0, 1, false).
			AddItem(primitive, height, 1, true).
			AddItem(nil, 0, 1, false), width, 1, true).
		AddItem(nil, 0, 1, false)
}

func rawOutput(primitive tview.Primitive) *tview.Flex {
	return tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(nil, 0, 1, false).
			AddItem(primitive, 1, 1, true).
			AddItem(nil, 0, 1, false), 1, 1, true)
}

func failWindow(reason string) *tview.Modal {
	fail := tview.NewModal().
		SetText(reason).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			closePopup()
		})

	fail.SetBackgroundColor(tcell.ColorHotPink)
	return fail
}

func successWindow(message string) *tview.Modal {
	wnd := tview.NewModal().
		SetText(message).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			// update info window & mainFrame window
			info = infoWindow()
			mainFrame = mainFrameWindow()
			updateView()
			root.SwitchToPage(pageMain)
		})

	wnd.SetBackgroundColor(tcell.ColorDarkGreen)
	return wnd
}

func passwordPrompt(path string) *tview.Flex {
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
			addAndShowPopup("passwordPrompt", failWindow("Failed Storing Master Key!!!"))
		} else {
			masterKey.path = path
			addAndShowPopup("passwordPrompt", successWindow(fmt.Sprint("Successfully Stored Master Key!!!\n", path)))
		}
	})
	form.SetFocus(0)

	return modal(40, 10, form)
}

func keyGenWindow() (content *tview.Flex) {
	text := tview.NewTextView().
		SetTextAlign(tview.AlignLeft).
		SetDynamicColors(true)

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
		addAndShowPopup("keygen", passwordPrompt(inputField.GetText()))
	})
	form.AddButton("Cancel", nil)
	form.SetFocus(0)

	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow).
		SetBorder(true).
		SetTitle(keyGenWindowTitle)
	flex.AddItem(text, 0, 1, false)
	flex.AddItem(form, 0, 1, true)

	return flex
}
