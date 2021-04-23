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

func showFailWindow(title string, msg string) {
	fail := tview.NewModal().
		SetText(msg).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			root.HidePage("fail")
		})

	fail.SetTitle(title)
	fail.SetBackgroundColor(tcell.ColorHotPink)

	root.AddPage("fail", fail, true, true)
}

func showSuccessWindow(title string, msg string, callback func()) {
	succ := tview.NewModal().
		SetText(msg).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			// update info window & mainFrame window
			root.RemovePage("success")
			if callback != nil {
				callback()
			}
		})

	succ.SetTitle(title)
	succ.SetBackgroundColor(tcell.ColorDarkGreen)

	root.AddPage("success", succ, true, true)
}
