package main

import (
	"fmt"
	"time"

	"github.com/rivo/tview"
)

var infoString = `
MASTER KEY CREATED:
[green]%v

[-:-:-]
NUM LABLED KEYS:
[green]%v
`

var keyLoadedString = `
Use ArrowKeys ← ↑ → ↓ To Select Keys

Derived keys will be uncovered when selected.
`

func infoWindow() (content *tview.Flex) {
	info = tview.NewFlex()
	info.SetDirection(tview.FlexRow)
	info.SetBorder(true)
	refreshInfo()
	return info
}

func refreshInfo() {
	info.Clear()
	info.AddItem(infoText(), 0, 8, false)

	// scroll to end to align text to bottom
	keyloadtv := keyLoadedText()
	info.AddItem(keyloadtv, 0, 2, false)
	keyloadtv.ScrollToEnd()
}

func infoText() *tview.TextView {
	textview := tview.NewTextView()
	textview.SetDynamicColors(true)

	if masterKey != nil {
		fmt.Fprintf(textview, infoString,
			time.Unix(masterKey.createdAt, 0).Local().Format(time.RFC822),
			len(masterKey.lables),
		)
	}
	return textview
}

func keyLoadedText() *tview.TextView {
	textview := tview.NewTextView()
	textview.SetDynamicColors(true)
	if masterKey != nil {
		fmt.Fprint(textview, "\n\n\n\n\n\n\n\n\n")
		fmt.Fprint(textview, keyLoadedString)
	}
	return textview
}
