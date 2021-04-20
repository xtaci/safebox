package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func MainWindow() (content tview.Primitive) {
	// Create a frame for the subtitle and navigation infos.
	box := tview.NewBox().
		SetBorder(true).
		SetBackgroundColor(tcell.ColorDimGray)

	frame := tview.NewFrame(box).
		SetBorders(0, 0, 0, 0, 0, 0).
		AddText("not loaded", true, tview.AlignCenter, tcell.ColorWhite)
	return frame
}

func FooterNotLoaded() (content tview.Primitive) {
	// The bottom row has some info on where we are.
	info := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false)

	// only F1, F2 Esc enabled
	keys := []tcell.Key{tcell.KeyF1, tcell.KeyF2}
	for _, key := range keys {
		fmt.Fprintf(info, `[%s] ["%s"][darkcyan]%s[white][""]  `, keyNames[key], keyNames[key], shortCuts[key])
	}
	return info
}
