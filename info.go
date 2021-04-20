package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var KeyUsage = `
how to use
`

func Info() (content tview.Primitive) {
	box := tview.NewBox().
		SetBorder(true)

	frame := tview.NewFrame(box).
		SetBorders(0, 0, 0, 0, 0, 0).
		AddText(KeyUsage, true, tview.AlignCenter, tcell.ColorWhite)
	return frame
}
