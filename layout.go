package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type windowState byte

const (
	normalWindow windowState = iota
	shortCutsActivated
)

var (
	root       *tview.Pages
	background *tview.TextView
	mainFrame  *tview.Flex
	body       *tview.Flex
	info       *tview.Flex
	footer     tview.Primitive
	layout     *tview.Flex
	state      windowState
)

// global shortcuts handling
func globalInputCapture(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyF1:
		mainFrame = keyGenWindow()
		updateView()
		state = shortCutsActivated
		return nil
	case tcell.KeyF2:
		return nil
	case tcell.KeyESC:
		if state == shortCutsActivated {
			body.RemoveItem(mainFrame)
			mainFrame = mainFrameWindow()
			body.AddItem(mainFrame, 0, 80, true)
			state = normalWindow
			return nil
		} else {
			app.Stop()
			return event
		}
	default:
		return event
	}
}

func initLayouts() {
	// Main frame
	mainFrame = mainFrameWindow()
	info = infoWindow()
	footer = footerWindow()
	body = tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(info, 0, 20, false).
		AddItem(mainFrame, 0, 80, true)

	// Create the layout.
	layout = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(body, 0, 1, true).
		AddItem(footer, 1, 1, false)

	root = tview.NewPages().
		AddPage("main", layout, true, true)

	state = normalWindow

	// Input capture
	app.SetInputCapture(globalInputCapture)
}

func updateView() {
	body = tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(info, 0, 20, false).
		AddItem(mainFrame, 0, 80, true)

	// Create the layout.
	layout = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(body, 0, 1, true).
		AddItem(footer, 1, 1, false)

	root = tview.NewPages().
		AddPage("main", layout, true, true)

	app.SetRoot(root, true)
}
