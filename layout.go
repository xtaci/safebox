package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var mainFrame *tview.Flex

// initial page
func layoutInit() tview.Primitive {
	// Main frame
	mainFrame = mainFrameMasterKeyNotLoaded()
	body := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(infoNotLoaded(), 0, 20, false).
		AddItem(mainFrame, 0, 80, true)

	// Create the layout.
	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(body, 0, 1, true).
		AddItem(footerNotLoaded(), 1, 1, false)

	// Input capture
	mainFrame.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyF1 {
			body.RemoveItem(mainFrame)
			mainFrame = keyGenWindow()
			body.AddItem(mainFrame, 0, 80, true)
			return nil
		}
		return event
	})

	return layout
}

// layout when master keys confirmed
func layoutLoaded() tview.Primitive {
	// Main frame
	mainFrame = tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(infoNotLoaded(), 0, 20, false).
		AddItem(mainFrameMasterKeyNotLoaded(), 0, 80, true)

	// Create the layout.
	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(mainFrame, 0, 1, true).
		AddItem(footerNotLoaded(), 1, 1, false)

	return layout
}
