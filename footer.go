package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func footerWindow() *tview.TextView {
	// The bottom row has some info on where we are.
	info := tview.NewTextView().
		SetToggleHighlights(true).
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false)

	//keys := []tcell.Key{tcell.KeyF1, tcell.KeyF2, tcell.KeyEsc}
	keys := []tcell.Key{tcell.KeyF1, tcell.KeyF2, tcell.KeyEsc, tcell.KeyCtrlC}
	for _, key := range keys {
		fmt.Fprintf(info, `%v ["%d"][darkcyan]%s[white][""] `, tview.Escape("["+keyNames[key]+"]"), key, shortCuts[key])
	}
	return info
}
