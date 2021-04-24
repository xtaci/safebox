package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rivo/tview"
)

var infoString = `
Master Key SHA256:
[green]%v

[-:-:-]Master Keys Created At:
[green]%v

[-:-:-]Number of keys with label:
[green]%v
`

var instructionsString = `
Instructions

1) Use ArrowKeys ← ↑ → ↓ To Select Keys, masks on derived keys will be uncovered when selected.

2) Press <Enter> on 'DERIVED KEY' column to export.

3) Press <Enter> on 'NAME' column to set label.
`

func infoWindow() (content *tview.Flex) {
	layoutInfo = tview.NewFlex()
	layoutInfo.SetDirection(tview.FlexRow)
	layoutInfo.SetBorder(true)
	refreshInfo()
	return layoutInfo
}

func refreshInfo() {
	layoutInfo.Clear()
	layoutInfo.AddItem(infoMasterKey(), 0, 1, false)

	// scroll to end to align text to bottom
	operations := infoInstructions()
	layoutInfo.AddItem(operations, 0, 1, false)
	operations.ScrollToEnd()
}

// master key info textview
func infoMasterKey() *tview.TextView {
	textview := tview.NewTextView()
	textview.SetDynamicColors(true)
	textview.SetWrap(true)

	if masterKey != nil {
		md := sha256.Sum256(masterKey.masterKey[:])
		fmt.Fprintf(textview, infoString,
			hexutil.Encode(md[:]),
			time.Unix(masterKey.createdAt, 0).Local().Format(time.RFC822),
			len(masterKey.labels),
		)
	}
	return textview
}

// instructions info textview
func infoInstructions() *tview.TextView {
	textview := tview.NewTextView()
	textview.SetDynamicColors(true)
	textview.SetWrap(true)
	textview.SetWordWrap(true)
	if masterKey != nil {
		fmt.Fprint(textview, strings.Repeat("\n", 100))
		fmt.Fprint(textview, instructionsString)
	}
	return textview
}
