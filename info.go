package main

import (
	"fmt"
	"time"

	"github.com/rivo/tview"
)

var infoString = `
MASTER KEY CREATED:%v
NUM LABLED KEYS:%v
`

func infoWindow() (content *tview.Flex) {
	info := tview.NewTextView()
	if masterKey != nil {
		fmt.Fprintf(info, infoString,
			time.Unix(masterKey.createdAt, 0),
			len(masterKey.labels),
		)
	}

	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow).
		AddItem(info, 0, 8, false).
		SetTitle("- KEY INFO -").
		SetBorder(true)

	return flex
}
