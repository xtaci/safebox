package main

import "github.com/rivo/tview"

// initial page
func layoutInit() tview.Primitive {
	// Main frame
	mainFrame := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(mainFrameNotLoaded(), 0, 80, true).
		AddItem(Info(), 0, 20, true)

	// Create the layout.
	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(mainFrame, 0, 1, true).
		AddItem(footerNotLoaded(), 1, 1, false)

	return layout
}

// layout when master keys confirmed
func layoutLoaded() tview.Primitive {
	// Main frame
	mainFrame := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(mainFrameNotLoaded(), 0, 80, true).
		AddItem(Info(), 0, 20, true)

	// Create the layout.
	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(mainFrame, 0, 1, true).
		AddItem(footerNotLoaded(), 1, 1, false)

	return layout
}
