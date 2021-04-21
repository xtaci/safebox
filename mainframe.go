package main

import (
	"fmt"

	"github.com/rivo/tview"
)

var verString = "VER 1.0"
var mainFrameTitle = fmt.Sprintf("- SAFEBOX KEY MANGEMENT SYSTEM %v -", verString)

// displays when master key is not loaded
func mainFrameMasterKeyNotLoaded() (content *tview.Flex) {
	text := tview.NewTextView()
	text.SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter)
	fmt.Fprintf(text, `[red]KEY NOT LOADED
PLEASE LOAD A MASTER KEY[yellow][F2][red] OR GENERATE ONE[yellow][F1][red] FIRST`)

	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow).
		SetBorder(true).
		SetTitle(mainFrameTitle)

	flex.AddItem(tview.NewBox(), 0, 8, false)
	flex.AddItem(text, 0, 1, true)
	flex.AddItem(tview.NewBox(), 0, 8, false)
	return flex
}

// displays when a master key is loaded
func mainFrameMasterKeyLoaded() (content *tview.Flex) {
	text := tview.NewTextView()
	text.SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter)

	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow).
		SetBorder(true).
		SetTitle(mainFrameTitle)

	flex.AddItem(tview.NewBox(), 0, 8, false)
	flex.AddItem(text, 0, 1, false)
	flex.AddItem(tview.NewBox(), 0, 8, false)
	return flex
}
