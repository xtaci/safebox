package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type windowState byte

const (
	wndNotLoaded windowState = iota
	wndKeyGen
)

var (
	pages      *tview.Pages
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
		body.RemoveItem(mainFrame)
		mainFrame = keyGenWindow()
		body.AddItem(mainFrame, 0, 80, true)
		state = wndKeyGen
		return nil
	case tcell.KeyESC:
		if state == wndKeyGen {
			body.RemoveItem(mainFrame)
			mainFrame = mainFrameWindow()
			body.AddItem(mainFrame, 0, 80, true)
			state = wndNotLoaded
			return nil
		}
		return event
	default:
		return event
	}
}

// container
func container() *tview.Pages {
	// Main frame
	mainFrame = mainFrameWindow()
	info = infoNotLoaded()
	footer = footerNotLoaded()
	body = tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(info, 0, 20, false).
		AddItem(mainFrame, 0, 80, true)

	// Create the layout.
	layout = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(body, 0, 1, true).
		AddItem(footer, 1, 1, false)

	pages = tview.NewPages().
		AddPage("main", layout, true, true)

	state = wndNotLoaded

	// Input capture
	app.SetInputCapture(globalInputCapture)

	return pages
}
