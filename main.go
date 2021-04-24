package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// The application.
var app = tview.NewApplication()

// The master key
var masterKey *MasterKey

var (
	keyNames = map[tcell.Key]string{
		tcell.KeyF1:    "F1",
		tcell.KeyF2:    "F2",
		tcell.KeyF3:    "F3",
		tcell.KeyF4:    "F4",
		tcell.KeyEsc:   "ESC",
		tcell.KeyCtrlC: "Ctrl-C",
	}

	shortCuts = map[tcell.Key]string{
		tcell.KeyF1:    "GENERATE MASTER KEY",
		tcell.KeyF2:    "LOAD MASTER KEY",
		tcell.KeyEsc:   "BACK",
		tcell.KeyCtrlC: "QUIT",
	}
)

var theme = tview.Theme{
	PrimitiveBackgroundColor:    tcell.ColorWhite,
	ContrastBackgroundColor:     tcell.ColorBlue,
	MoreContrastBackgroundColor: tcell.ColorGreen,
	BorderColor:                 tcell.ColorBlack,
	TitleColor:                  tcell.ColorRed,
	GraphicsColor:               tcell.ColorBlack,
	PrimaryTextColor:            tcell.ColorBlack,
	SecondaryTextColor:          tcell.ColorBlack,
	TertiaryTextColor:           tcell.ColorGreen,
	InverseTextColor:            tcell.ColorBlue,
	ContrastSecondaryTextColor:  tcell.ColorDarkCyan,
}

func main() {
	tview.Styles = theme
	initLayouts()
	// Start the application and set root to Cover
	if err := app.SetRoot(layoutRoot, true).Run(); err != nil {
		panic(err)
	}
}
