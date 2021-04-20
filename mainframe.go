package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func mainFrameNotLoaded() (content tview.Primitive) {
	// Create a frame for the subtitle and navigation infos.
	box := tview.NewBox().
		SetBorder(true).
		SetBackgroundColor(tcell.ColorDimGray)

	frame := tview.NewFrame(box).
		SetBorders(0, 0, 0, 0, 0, 0).
		AddText("not loaded", true, tview.AlignCenter, tcell.ColorWhite)
	return frame
}
