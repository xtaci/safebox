package main

import (
	"github.com/rivo/tview"
)

func showChangePasswordWindow() {
	const (
		windowName   = "changePassword"
		windowWidth  = 40
		windowHeight = 10
		windowTitle  = "CHANGE MASTERKEY PASSWORD 🔑"
	)

	form := tview.NewForm()
	form.SetTitle(windowTitle)
	form.SetBorder(true)
	passwordField := tview.NewInputField().SetLabel("Password").
		SetFieldWidth(64).
		SetMaskCharacter('*')

	passwordFieldConfirm := tview.NewInputField().SetLabel("Password Confirm").
		SetFieldWidth(64).
		SetMaskCharacter('*')

	form.AddButton("OK", func() {
		if passwordField.GetText() != passwordFieldConfirm.GetText() {
			showFailWindow("PASSWORD MISMATCH")
		} else {
			masterKey.changePassword([]byte(passwordField.GetText()))
			masterKey.store(masterKey.path)
			showSuccessWindow("MASTER KEY PASSWORD CHANGED!!!", func() {
				layoutRoot.RemovePage(windowName)
			})
		}
	})

	form.AddFormItem(passwordField)
	form.AddFormItem(passwordFieldConfirm)
	form.SetFocus(0)

	layoutRoot.AddPage(windowName, popup(windowWidth, windowHeight, form), true, true)
}
