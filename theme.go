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
	PrimaryTextColor:            tcell.ColorWhite,
	SecondaryTextColor:          tcell.ColorAntiqueWhite,
	TertiaryTextColor:           tcell.ColorWhite,
	InverseTextColor:            tcell.ColorBlack,
	ContrastSecondaryTextColor:  tcell.ColorDarkCyan,
}
