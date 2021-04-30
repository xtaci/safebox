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
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			layoutRoot.RemovePage("fail")
		})

	fail.SetBackgroundColor(tcell.ColorHotPink)

	layoutRoot.AddPage("fail", fail, true, true)
}

func showSuccessWindow(msg string, callback func()) {
	succ := tview.NewModal().
		SetText(msg).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			// update info window & mainFrame window
			layoutRoot.RemovePage("success")
			if callback != nil {
				callback()
			}
		})

	succ.SetBackgroundColor(tcell.ColorLightGreen)

	layoutRoot.AddPage("success", succ, true, true)
}

func showInfoWindow(msg string, callback func()) {
	succ := tview.NewModal().
		SetText(msg).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			// update info window & mainFrame window
			layoutRoot.RemovePage("info")
			if callback != nil {
				callback()
			}
		})

	succ.SetBackgroundColor(tcell.ColorDarkOrange)

	layoutRoot.AddPage("info", succ, true, true)
}
