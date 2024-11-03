package main

import (
	"github.com/rivo/tview"
)

func showChangePasswordWindow() {
	const (
		windowName   = "changePassword"
		windowWidth  = 40
		windowHeight = 10
		windowTitle  = S_WINDOW_CHANGEPASS_TITLE
	)

	form := tview.NewForm()
	form.SetTitle(windowTitle)
	form.SetBorder(true)
	passwordField := tview.NewInputField().SetLabel(S_WINDOW_CHANGEPASS_LABEL_PASSWORD).
		SetFieldWidth(0).
		SetMaskCharacter('*')

	passwordFieldConfirm := tview.NewInputField().SetLabel(S_WINDOW_CHANGEPASS_LABEL_CONFIRM).
		SetFieldWidth(0).
		SetMaskCharacter('*')

	form.AddButton(S_WINDOW_CHANGEPASS_BUTTON_OK, func() {
		if passwordField.GetText() != passwordFieldConfirm.GetText() {
			showFailWindow(S_MSG_PASSWORD_MISMATCH)
		} else {
			masterKey.changePassword([]byte(passwordField.GetText()))
			masterKey.store(masterKey.path)
			showSuccessWindow(S_MSG_PASSWORD_CHANGED, func() {
				layoutRoot.RemovePage(windowName)
			})
		}
	})

	form.AddFormItem(passwordField)
	form.AddFormItem(passwordFieldConfirm)
	form.SetFocus(0)

	layoutRoot.AddPage(windowName, popup(windowWidth, windowHeight, form), true, true)
}
