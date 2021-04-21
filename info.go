package main

import (
	"fmt"

	"github.com/rivo/tview"
)

var infoString = `
MASTER KEY CREATED:\t%v
NUM DERIVED KEYS:\t%v

SYSTEM INFORMATION:
OS: \t%v
KERNEL:\t%v

COPYRIGHT 2021 (C) xtaci
`

func infoNotLoaded() (content *tview.Flex) {
	info := tview.NewTextView()
	fmt.Fprintf(info, infoString)

	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow).
		AddItem(info, 0, 8, false).
		SetTitle("- KEY INFO -").
		SetBorder(true)

	return flex
}
