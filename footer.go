package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var verString = "//safebox //version 1.0 "

func footerWindow() *tview.Flex {
	// The bottom row has some info on where we are.
	layoutFooter = tview.NewFlex()
	refreshFooter()
	return layoutFooter
}

func refreshFooter() {
	layoutFooter.Clear()

	layoutShortcuts = tview.NewTextView().
		SetToggleHighlights(true).
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false)

		// shortcuts
	if masterKey == nil {
		keys := []tcell.Key{tcell.KeyCtrlG, tcell.KeyCtrlL, tcell.KeyEsc, tcell.KeyCtrlC}
		for _, key := range keys {
			fmt.Fprintf(layoutShortcuts, `[darkorange]%v ["%d"][black]%s[""] `, tview.Escape("<"+keyNames[key]+">"), key, shortCuts[key])
		}
	} else {
		keys := []tcell.Key{tcell.KeyCtrlG, tcell.KeyCtrlL, tcell.KeyCtrlP, tcell.KeyEsc, tcell.KeyCtrlC}
		for _, key := range keys {
			fmt.Fprintf(layoutShortcuts, `[darkorange]%v ["%d"][black]%s[""] `, tview.Escape("<"+keyNames[key]+">"), key, shortCuts[key])
		}
	}

	// right aligned version string
	versionTextView := tview.NewTextView().
		SetTextAlign(tview.AlignRight).
		SetDynamicColors(true).
		SetWrap(false)

	fmt.Fprint(versionTextView, verString)

	//  flex
	layoutFooter.AddItem(layoutShortcuts, 0, 8, false)
	layoutFooter.AddItem(versionTextView, 0, 2, false)
}
