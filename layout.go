package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type windowState byte

const (
	pageCover   = "COVER"
	pageMain    = "MAIN"
	pageKeyGen  = "KEYGEN"
	pageKeyLoad = "KEYLOAD"
)

type popupWindow struct {
	name      string
	primitive tview.Primitive
}

var (
	root       *tview.Pages
	cover      *tview.Flex
	background *tview.TextView
	mainFrame  *tview.Flex
	info       *tview.Flex
	footer     tview.Primitive
	popupStack []popupWindow // popup window stack
)

func addAndShowPopup(name string, primitive tview.Primitive) {
	root.AddAndSwitchToPage(name, primitive, true)
	popupStack = append(popupStack, popupWindow{name, primitive})
}
func closePopup() {
	if len(popupStack) > 0 {
		popupStack = popupStack[:len(popupStack)-1]
		updateView()
	}
}

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
		if name, _ := root.GetFrontPage(); name != pageKeyGen {
			primitive := keyGenWindow()
			addAndShowPopup(pageKeyGen, primitive)
		}
		return nil
	case tcell.KeyF2:
		if name, _ := root.GetFrontPage(); name != pageKeyGen {
			primitive := loadKeyWindow()
			addAndShowPopup(pageKeyLoad, primitive)
		}
		return nil
	case tcell.KeyESC:
		if len(popupStack) > 0 {
			closePopup()
		} else {
			app.Stop()
		}
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
	body := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(info, 0, 20, false).
		AddItem(mainFrame, 0, 80, true)

	// Create the layout.
	layout := tview.NewFlex().
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

func updateView() {
	body := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(info, 0, 20, false).
		AddItem(mainFrame, 0, 80, true)

	// Create the layout.
	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(body, 0, 1, true).
		AddItem(footer, 1, 1, false)

	root = tview.NewPages().
		AddPage(pageMain, layout, true, true)

	app.SetRoot(root, true)
}
