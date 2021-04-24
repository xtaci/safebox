package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func footerWindow() *tview.TextView {
	// The bottom row has some info on where we are.
	layoutFooter = tview.NewTextView().
		SetToggleHighlights(true).
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false)

	refreshFooter()
	return layoutFooter
}

func refreshFooter() {
	layoutFooter.Clear()
	if masterKey == nil {
		keys := []tcell.Key{tcell.KeyF1, tcell.KeyF2, tcell.KeyEsc, tcell.KeyCtrlC}
		for _, key := range keys {
			fmt.Fprintf(layoutFooter, `[darkorange]%v ["%d"][black]%s[""] `, tview.Escape("<"+keyNames[key]+">"), key, shortCuts[key])
		}
	} else {
		keys := []tcell.Key{tcell.KeyF1, tcell.KeyF2, tcell.KeyF3, tcell.KeyEsc, tcell.KeyCtrlC}
		for _, key := range keys {
			fmt.Fprintf(layoutFooter, `[darkorange]%v ["%d"][black]%s[""] `, tview.Escape("<"+keyNames[key]+">"), key, shortCuts[key])
		}
	}
}
