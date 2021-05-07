package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

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
	//                      |_______layoutShortCuts
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
	layoutFooter        *tview.Flex
	layoutShortcuts     *tview.TextView
)

var (
	keyNames = map[tcell.Key]string{
		tcell.KeyF1:    "F1",
		tcell.KeyF2:    "F2",
		tcell.KeyF3:    "F3",
		tcell.KeyF4:    "F4",
		tcell.KeyEsc:   "ESC",
		tcell.KeyCtrlG: "F1|Ctrl-G",
		tcell.KeyCtrlL: "F2|Ctrl-L",
		tcell.KeyCtrlP: "F3|Ctrl-P",
		tcell.KeyCtrlC: "Ctrl-C",
	}

	shortCuts = map[tcell.Key]string{
		tcell.KeyCtrlG: "Generate master key",
		tcell.KeyCtrlL: "Load master key",
		tcell.KeyCtrlP: "Change password",
		tcell.KeyEsc:   "Back",
		tcell.KeyCtrlC: "Quit",
	}
)

// global function keys handling
func globalInputCapture(event *tcell.EventKey) *tcell.EventKey {
	// check cover page
	name, _ := layoutRoot.GetFrontPage()
	if name == pageCover {
		layoutRoot.SwitchToPage(pageMain)
		return nil
	}

	// non cover page
	switch event.Key() {
	case tcell.KeyF1, tcell.KeyCtrlG:
		if name, _ := layoutRoot.GetFrontPage(); name == pageMain {
			showKeyEntropyInputWindow()
			layoutShortcuts.Highlight(fmt.Sprint(tcell.KeyF1))
		}
		return nil

	case tcell.KeyF2, tcell.KeyCtrlL:
		if name, _ := layoutRoot.GetFrontPage(); name == pageMain {
			showLoadKeyWindow()
			layoutShortcuts.Highlight(fmt.Sprint(tcell.KeyF2))
		}
		return nil

	case tcell.KeyF3, tcell.KeyCtrlP:
		if name, _ := layoutRoot.GetFrontPage(); name == pageMain {
			showChangePasswordWindow()
			layoutShortcuts.Highlight(fmt.Sprint(tcell.KeyF3))
		}
		return nil

	case tcell.KeyEsc:
		if name, _ := layoutRoot.GetFrontPage(); name != pageMain {
			layoutRoot.RemovePage(name)
			if name, _ := layoutRoot.GetFrontPage(); name == pageMain {
				layoutShortcuts.Highlight(layoutShortcuts.GetHighlights()...)
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

func refresh() {
	refreshFooter()
	refreshInfo()
	refreshMainFrame()
}
