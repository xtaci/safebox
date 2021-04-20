package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func footerNotLoaded() (content tview.Primitive) {
	// The bottom row has some info on where we are.
	info := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false)

	// only F1, F2 Esc enabled
	keys := []tcell.Key{tcell.KeyF1, tcell.KeyF2}
	for _, key := range keys {
		fmt.Fprintf(info, `[%s] ["%s"][darkcyan]%s[white][""]  `, keyNames[key], keyNames[key], shortCuts[key])
	}
	return info
}
