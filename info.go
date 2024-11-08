package main

import (
	"crypto/sha256"
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var infoString = `
[-:-:-]Version
[darkblue]%v

[-:-:-]Location:
[darkblue]%v

[-:-:-]Master Key SHA256:
[darkblue]%v

[-:-:-]Master Keys Created At:
[darkblue]%v

[-:-:-]Number of keys with label:
[darkblue]%v

[-:-:-]System:
[darkblue]%v %v

`

var instructionsString = `
Instructions

1) Use ArrowKeys [darkred]← ↑ → ↓ [-:-:-]To Select Keys, masks on derived keys will be uncovered when selected.

2) Press [darkred]<Enter>[-:-:-] on 'Derived Keys' column to export.

3) Press [darkred]<Enter>[-:-:-] on 'Label' column to set label.

4) Use [darkred]<TAB>[-:-:-] to focus on different items.
`

func infoWindow() (content *tview.Flex) {
	layoutInfo = tview.NewFlex()
	layoutInfo.SetDirection(tview.FlexRow)
	layoutInfo.SetBorder(true)
	layoutInfo.SetTitle(S_MAIN_FRAME_TITLE)
	layoutInfo.SetTitleColor(tcell.ColorWhite)
	layoutInfo.SetBorderColor(tcell.ColorWhite)
	layoutInfo.SetBackgroundColor(tcell.ColorBlue)
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
	textview.SetBackgroundColor(tcell.ColorBlue)
	textview.SetWrap(true)

	if masterKey != nil {
		md := sha256.Sum256(masterKey.masterKey[:])
		fmt.Fprintf(textview, infoString,
			VERSION,
			masterKey.path,
			hexutil.Encode(md[:]),
			time.Unix(masterKey.createdAt, 0).Local().Format(time.RFC822),
			len(masterKey.labels),
			runtime.GOOS,
			runtime.GOARCH,
		)
	} else {
		fmt.Fprintf(textview, infoString,
			VERSION,
			"N/A",
			"N/A",
			"N/A",
			"N/A",
			runtime.GOOS,
			runtime.GOARCH,
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
	textview.SetBackgroundColor(tcell.ColorBlue)
	fmt.Fprint(textview, strings.Repeat("\n", 100))
	fmt.Fprint(textview, instructionsString)
	return textview
}
