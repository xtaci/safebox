package main

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Cover returns the cover page.
func Cover() (content tview.Primitive) {
	// What's the size of the logo?
	lines := strings.Split(logo, "\n")
	logoWidth := 0
	logoHeight := len(lines)
	for _, line := range lines {
		if len(line) > logoWidth {
			logoWidth = len(line)
		}
	}
	logoBox := tview.NewTextView().
		SetTextColor(tcell.ColorGreen)
	fmt.Fprint(logoBox, logo)

	// Create a frame for the subtitle and navigation infos.
	frame := tview.NewFrame(tview.NewBox()).
		SetBorders(0, 0, 0, 0, 0, 0).
		AddText(subtitle, true, tview.AlignCenter, tcell.ColorWhite).
		AddText("", true, tview.AlignCenter, tcell.ColorWhite).
		AddText(navigation, true, tview.AlignCenter, tcell.ColorDarkMagenta)

	// Create a Flex layout that centers the logo and subtitle.
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tview.NewBox(), 0, 7, false).
		AddItem(tview.NewFlex().
			AddItem(tview.NewBox(), 0, 1, false).
			AddItem(logoBox, logoWidth, 1, true).
			AddItem(tview.NewBox(), 0, 1, false), logoHeight, 1, true).
		AddItem(frame, 0, 10, false)

	// Input capture
	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		initBody := layoutInit()
		app.SetRoot(initBody, true)
		return event
	})

	return flex
}