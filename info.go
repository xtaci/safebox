package main

import (
	"fmt"
	"time"

	"github.com/rivo/tview"
)

var infoString = `
[lightgray]MASTER KEY CREATED:
%v

NUM LABLED KEYS:
%v
`

var keyLoadedString = `
Use ArrowKeys ← ↑ → ↓ To Select Keys

Derived keys will be uncovered when selected.
`

func infoWindow() (content *tview.Flex) {
	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow).
		AddItem(infoText(), 0, 1, false).
		SetTitle("- KEY INFO -").
		SetBorder(true)

	return flex
}

func refreshInfo() {
	info.Clear()
	info.AddItem(infoText(), 0, 1, false)
}

func infoText() *tview.TextView {
	textview := tview.NewTextView()
	textview.SetDynamicColors(true)

	if masterKey != nil {
		fmt.Fprintf(textview, infoString,
			time.Unix(masterKey.createdAt, 0).Local().Format(time.RFC822),
			len(masterKey.lables),
		)

		fmt.Fprintf(textview, keyLoadedString)
	}
	return textview
}
