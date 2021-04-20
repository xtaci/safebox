package main

import (
	"fmt"

	"github.com/rivo/tview"
)

var verString = "VER 1.0"

func mainFrameNotLoaded() (content tview.Primitive) {
	info := tview.NewBox().SetBorder(true).SetTitle(fmt.Sprintf("- SAFEBOX KEY MANGEMENT SYSTEM %v -", verString))

	text := tview.NewTextView().
		SetTextAlign(tview.AlignCenter)

	fmt.Fprintf(text, "KEY NOT LOADED")

	flex := tview.NewFrame(info)
	return flex
}
