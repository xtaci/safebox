package main

import (
	"github.com/awnumar/memguard"
	"github.com/rivo/tview"
)

var VERSION = "undefined"

// The application.
var app *tview.Application

// The master key
var masterKey *MasterKey

func main() {
	fixCharset()
	// Safely terminate in case of an interrupt signal
	memguard.CatchInterrupt()

	// Purge the session when we return
	defer memguard.Purge()

	app = tview.NewApplication()
	tview.Styles = theme
	initLayouts()
	// Start the application and set root to Cover
	if err := app.SetRoot(layoutRoot, true).Run(); err != nil {
		panic(err)
	}
}
