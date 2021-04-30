package main

import (
	"os"
	"os/exec"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// The application.
var app *tview.Application

// The master key
var masterKey *MasterKey

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

var (
	wideCharset = []string{"zh_", "jp_", "ko_", "ja_", "th_", "hi_"}
)

func main() {
	locale := os.Getenv("LANG")

	var asianCharset bool
	for k := range wideCharset {
		if strings.HasPrefix(locale, wideCharset[k]) {
			asianCharset = true
		}
	}

	if asianCharset {
		os.Setenv("LANG", "C.UTF-8")
		cmd := exec.Command(os.Args[0])
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
		return
	}

	app = tview.NewApplication()
	tview.Styles = theme
	initLayouts()
	// Start the application and set root to Cover
	if err := app.SetRoot(layoutRoot, true).Run(); err != nil {
		panic(err)
	}
}
