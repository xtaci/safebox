// Demo code for the List primitive.
package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const logo = `

   __|    \    __| __|  _ )   _ \ \ \  / 
 \__ \   _ \   _|  _|   _ \  (   | >  <  
 ____/ _/  _\ _|  ___| ___/ \___/  _/\_\ 
`

// The application.
var app = tview.NewApplication()

const (
	subtitle   = `safebox - UNIFIED KEY MANAGEMENT SYSTEM`
	navigation = `Press any key to continue...`
)

var (
	keyNames = map[tcell.Key]string{
		tcell.KeyF1:  "F1",
		tcell.KeyF2:  "F2",
		tcell.KeyF3:  "F3",
		tcell.KeyF4:  "F4",
		tcell.KeyESC: "ESC",
	}

	shortCuts = map[tcell.Key]string{
		tcell.KeyF1:  "Gen Master Key",
		tcell.KeyF2:  "Load Master Key",
		tcell.KeyF3:  "Derive Key",
		tcell.KeyF4:  "Label Key",
		tcell.KeyESC: "QUIT",
	}
)

func main() {

	// Start the application and set root to Cover
	if err := app.SetRoot(Cover(), true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
