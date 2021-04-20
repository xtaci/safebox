// Demo code for the List primitive.
package main

import (
	"fmt"
	"strings"

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
	navigation = `Ctrl-C: Exit`
	mouse      = `(or use your mouse)`
)

var (
	keyNames = map[tcell.Key]string{
		tcell.KeyF1: "F1",
		tcell.KeyF2: "F2",
		tcell.KeyF3: "F3",
	}
)

// Cover returns the cover page.
func Cover() (title string, shortCut tcell.Key, content tview.Primitive) {
	// What's the size of the logo?
	lines := strings.Split(logo, "\n")
	logoWidth := 0
	logoHeight := len(lines)
	for _, line := range lines {
		if len(line) > logoWidth {
			logoWidth = len(line)
		}
	}
	logoBox := tview.NewTextView().
		SetTextColor(tcell.ColorGreen)
	fmt.Fprint(logoBox, logo)

	// Create a frame for the subtitle and navigation infos.
	frame := tview.NewFrame(tview.NewBox()).
		SetBorders(0, 0, 0, 0, 0, 0).
		AddText(subtitle, true, tview.AlignCenter, tcell.ColorWhite).
		AddText("", true, tview.AlignCenter, tcell.ColorWhite).
		AddText(navigation, true, tview.AlignCenter, tcell.ColorDarkMagenta).
		AddText(mouse, true, tview.AlignCenter, tcell.ColorDarkMagenta)

	// Create a Flex layout that centers the logo and subtitle.
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tview.NewBox(), 0, 7, false).
		AddItem(tview.NewFlex().
			AddItem(tview.NewBox(), 0, 1, false).
			AddItem(logoBox, logoWidth, 1, true).
			AddItem(tview.NewBox(), 0, 1, false), logoHeight, 1, true).
		AddItem(frame, 0, 10, false)

	return "Home", tcell.KeyF1, flex
}

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

type Slide func() (title string, shortCutKey tcell.Key, content tview.Primitive)

func main() {
	pages := tview.NewPages()
	slides := []Slide{
		Cover,
		MainKey,
	}

	// The bottom row has some info on where we are.
	info := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false).
		SetHighlightedFunc(func(added, removed, remaining []string) {
			pages.SwitchToPage(added[0])
		})

	for idx, slide := range slides {
		title, key, primitive := slide()
		pages.AddPage(title, primitive, true, idx == 0)
		fmt.Fprintf(info, `[%s] ["%s"][darkcyan]%s[white][""]  `, keyNames[key], keyNames[key], title)
	}

	// switch to home
	info.Highlight("Home")

	// Create the main layout.
	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(pages, 0, 1, true).
		AddItem(info, 1, 1, false)

	// shortcuts set
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		for _, slide := range slides {
			title, key, _ := slide()
			if event.Key() == key {
				info.Highlight(title)
				return nil
			}
		}
		return event
	})
	// Start the application.
	if err := app.SetRoot(layout, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
