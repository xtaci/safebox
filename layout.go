package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type windowState byte

const (
	pageBackground = "BACKGROUND"
	pageCover      = "COVER"
	pageMain       = "MAIN"
)

type popupWindow struct {
	name      string
	primitive tview.Primitive
}

var (
	root      *tview.Pages
	cover     *tview.Flex
	body      *tview.Flex
	layout    *tview.Flex
	mainFrame *tview.Flex
	info      *tview.Flex
	footer    *tview.TextView
)

// global shortcuts handling
func globalInputCapture(event *tcell.EventKey) *tcell.EventKey {
	// check cover page
	name, _ := root.GetFrontPage()
	if name == pageCover {
		root.SwitchToPage(pageMain)
		return nil
	}

	switch event.Key() {
	case tcell.KeyF1:
		if name, _ := root.GetFrontPage(); name == pageMain {
			showKeyGenWindow()
			footer.Highlight(fmt.Sprint(tcell.KeyF1))
		}
		return nil

	case tcell.KeyF2:
		if name, _ := root.GetFrontPage(); name == pageMain {
			showLoadKeyWindow()
			footer.Highlight(fmt.Sprint(tcell.KeyF2))
		}
		return nil

	case tcell.KeyEsc:
		if name, _ := root.GetFrontPage(); name != pageMain {
			root.HidePage(name)
		}
		return nil

	case tcell.KeyCtrlC:
		app.Stop()
		return nil
	default:
		return event
	}
}

// init pages
func initLayouts() {
	// Main frame
	mainFrame = mainFrameWindow()
	cover = coverPage()
	info = infoWindow()
	footer = footerWindow()
	body = tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(info, 0, 20, false).
		AddItem(mainFrame, 0, 80, true)

	body.SetBorderPadding(1, 1, 1, 1)

	// Create the layout.
	layout = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(body, 0, 1, true).
		AddItem(footer, 1, 1, false)

	root = tview.NewPages().
		AddPage(pageMain, layout, true, false).
		AddPage(pageCover, cover, true, true)

	root.SwitchToPage(pageCover)

	// Input capture
	app.SetInputCapture(globalInputCapture)
}

func refreshBody() {
	body.Clear()
	body.AddItem(info, 0, 20, false).
		AddItem(mainFrame, 0, 80, true)
}
