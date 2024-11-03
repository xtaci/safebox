package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var theme = tview.Theme{
	PrimitiveBackgroundColor:    tcell.ColorBlack,
	ContrastBackgroundColor:     tcell.ColorDarkGreen,
	MoreContrastBackgroundColor: tcell.ColorLightGreen,
	BorderColor:                 tcell.ColorDarkSeaGreen,
	TitleColor:                  tcell.ColorForestGreen,
	GraphicsColor:               tcell.ColorOrchid,
	PrimaryTextColor:            tcell.ColorLightGray,
	SecondaryTextColor:          tcell.ColorLightGreen,
	TertiaryTextColor:           tcell.ColorLightGreen,
	InverseTextColor:            tcell.ColorBlack,
	ContrastSecondaryTextColor:  tcell.ColorDarkCyan,
}
