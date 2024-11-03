package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func popup(width int, height int, primitive tview.Primitive) *tview.Flex {
	flex := tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(nil, 0, 1, false).
			AddItem(primitive, height, 1, true).
			AddItem(nil, 0, 1, false), width, 1, true).
		AddItem(nil, 0, 1, false)

	return flex
}

func showFailWindow(msg string) {
	fail := tview.NewModal().
		SetText(msg).
		AddButtons([]string{S_MODAL_BUTTON_OK}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			layoutRoot.RemovePage("fail")
		})

	fail.SetTitle(S_MODAL_TITLE_ERROR)
	fail.Box.SetBackgroundColor(tcell.ColorDarkRed)
	fail.SetBackgroundColor(tcell.ColorDarkRed)

	layoutRoot.AddPage("fail", fail, true, true)
}

func showSuccessWindow(msg string, callback func()) {
	succ := tview.NewModal().
		SetText(msg).
		AddButtons([]string{S_MODAL_BUTTON_OK}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			// update info window & mainFrame window
			layoutRoot.RemovePage("success")
			if callback != nil {
				callback()
			}
		})

	succ.SetTitle(S_MODAL_TITLE_SUCCESS)
	succ.Box.SetBackgroundColor(tcell.ColorDarkGreen)
	succ.SetBackgroundColor(tcell.ColorDarkGreen)
	layoutRoot.AddPage("success", succ, true, true)
}

func showInfoWindow(msg string, callback func()) {
	info := tview.NewModal().
		SetText(msg).
		AddButtons([]string{S_MODAL_BUTTON_OK}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			// update info window & mainFrame window
			layoutRoot.RemovePage("info")
			if callback != nil {
				callback()
			}
		})

	info.SetTitle(S_MODAL_TITLE_INFO)
	info.Box.SetBackgroundColor(tcell.ColorDarkBlue)
	info.SetBackgroundColor(tcell.ColorDarkBlue)
	layoutRoot.AddPage("info", info, true, true)
}
