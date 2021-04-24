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

	// structure of safebox windows
	// layoutRoot
	//    |______layoutCover
	//    |______layoutMain
	//              |______layoutFooter
	//              |______layoutBody
	//                      |______layoutMainBody
	//                      |            |______layoutMainBodyTable
	//                      |______layoutInfo

	layoutRoot          *tview.Pages
	layoutCover         *tview.Flex
	layoutBody          *tview.Flex
	layoutMain          *tview.Flex
	layoutMainBody      *tview.Flex
	layoutMainBodyTable *tview.Table
	layoutInfo          *tview.Flex
	layoutFooter        *tview.TextView
)

// global shortcuts handling
func globalInputCapture(event *tcell.EventKey) *tcell.EventKey {
	// check cover page
	name, _ := layoutRoot.GetFrontPage()
	if name == pageCover {
		layoutRoot.SwitchToPage(pageMain)
		return nil
	}

	switch event.Key() {
	case tcell.KeyF1:
		if name, _ := layoutRoot.GetFrontPage(); name == pageMain {
			showKeyGenWindow()
			layoutFooter.Highlight(fmt.Sprint(tcell.KeyF1))
		}
		return nil

	case tcell.KeyF2:
		if name, _ := layoutRoot.GetFrontPage(); name == pageMain {
			showLoadKeyWindow()
			layoutFooter.Highlight(fmt.Sprint(tcell.KeyF2))
		}
		return nil

	case tcell.KeyEsc:
		if name, _ := layoutRoot.GetFrontPage(); name != pageMain {
			layoutRoot.RemovePage(name)
			if name, _ := layoutRoot.GetFrontPage(); name == pageMain {
				layoutFooter.Highlight(layoutFooter.GetHighlights()...)
			}
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
	layoutMainBody = mainFrameWindow()
	layoutCover = coverPage()
	layoutInfo = infoWindow()
	layoutFooter = footerWindow()
	layoutBody = tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(layoutInfo, 0, 20, false).
		AddItem(layoutMainBody, 0, 80, true)

	layoutBody.SetBorderPadding(1, 1, 1, 1)

	// Create the layout.
	layoutMain = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(layoutBody, 0, 1, true).
		AddItem(layoutFooter, 1, 1, false)

	layoutRoot = tview.NewPages().
		AddPage(pageMain, layoutMain, true, false).
		AddPage(pageCover, layoutCover, true, true)

	layoutRoot.SwitchToPage(pageCover)

	// Input capture
	app.SetInputCapture(globalInputCapture)
}

func refreshBody() {
	layoutBody.Clear()
	layoutBody.AddItem(layoutInfo, 0, 20, false).
		AddItem(layoutMainBody, 0, 80, true)
}
