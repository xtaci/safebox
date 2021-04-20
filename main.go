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
	mouse      = `(or use your mouse)`
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

// Main Key generation
func MainKey() (title string, shortCut tcell.Key, content tview.Primitive) {
	// Create a Flex layout that centers the logo and subtitle.
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tview.NewBox(), 0, 7, false)

	return "Main Key", tcell.KeyF2, flex
}

// Key generation
func DeriveKey() (title string, shortCut tcell.Key, content tview.Primitive) {
	// Create a Flex layout that centers the logo and subtitle.
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tview.NewBox(), 0, 7, false)

	return "Drive Key", tcell.KeyF3, flex
}

func main() {
	// Main frame
	mainFrame := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(MainWindow(), 0, 80, true).
		AddItem(Info(), 0, 20, true)

	// Create the layout.
	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(mainFrame, 0, 1, true).
		AddItem(FooterNotLoaded(), 1, 1, false)

	splashShowed := false
	// capture any key
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if !splashShowed {
			app.SetRoot(layout, true)
			splashShowed = true
		}
		return event
	})

	// Start the application and set root to Cover
	if err := app.SetRoot(Cover(), true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
