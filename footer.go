package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

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
			fmt.Fprintf(layoutShortcuts, `[darkorange]%v ["%d"][white]%s[""] `, tview.Escape("<"+keyNames[key]+">"), key, shortCuts[key])
		}
	} else {
		keys := []tcell.Key{tcell.KeyCtrlG, tcell.KeyCtrlL, tcell.KeyCtrlP, tcell.KeyEsc, tcell.KeyCtrlC}
		for _, key := range keys {
			fmt.Fprintf(layoutShortcuts, `[darkorange]%v ["%d"][white]%s[""] `, tview.Escape("<"+keyNames[key]+">"), key, shortCuts[key])
		}
	}

	// right aligned version string
	versionTextView := tview.NewTextView().
		SetTextAlign(tview.AlignRight).
		SetDynamicColors(true).
		SetWrap(false)

	versionTextView.SetText(S_FOOTER_COPYRIGHT)

	//  flex
	layoutFooter.AddItem(layoutShortcuts, 0, 7, false)
	layoutFooter.AddItem(versionTextView, 0, 3, false)
}
