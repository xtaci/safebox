package main

import (
	"fmt"

	"github.com/rivo/tview"
)

var verString = "VER 1.0"

func mainFrameNotLoaded() (content tview.Primitive) {
	text := tview.NewTextView()
	text.SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter).
		SetBorderPadding(10, 10, 10, 10)

	fmt.Fprintf(text, `[red]KEY NOT LOADED
PLEASE LOAD A MASTER KEY[yellow][F2][red] OR GENERATE ONE[yellow][F1][red] FIRST`)

	flex := tview.NewFlex()
	flex.SetBorder(true).SetTitle(fmt.Sprintf("- SAFEBOX KEY MANGEMENT SYSTEM %v -", verString))
	flex.AddItem(text, 0, 1, false)
	return flex
}

func mainFrameLoaded() (content tview.Primitive) {
	return nil
}
