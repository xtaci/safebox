package main

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const logo = `
 ________  ________  ________ _______   ________  ________     ___    ___
|\   ____\|\   __  \|\  _____\\  ___ \ |\   __  \|\   __  \   |\  \  /  /|
\ \  \___|\ \  \|\  \ \  \__/\ \   __/|\ \  \|\ /\ \  \|\  \  \ \  \/  / /
 \ \_____  \ \   __  \ \   __\\ \  \_|/_\ \   __  \ \  \\\  \  \ \    / /
  \|____|\  \ \  \ \  \ \  \_| \ \  \_|\ \ \  \|\  \ \  \\\  \  /     \/
    ____\_\  \ \__\ \__\ \__\   \ \_______\ \_______\ \_______\/  /\   \
   |\_________\|__|\|__|\|__|    \|_______|\|_______|\|_______/__/ /\ __\
   \|_________|                                               |__|/ \|__|
`

const (
	subtitle   = `[::u]AN UNIFIED KEY MANAGEMENT SYSTEM`
	navigation = `[::bl]Press any key to continue...[-:-:-]`
)

// coverPage returns the coverPage page.
func coverPage() *tview.Flex {
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
		SetTextColor(tcell.ColorDarkCyan)
	fmt.Fprint(logoBox, logo)

	// Create a frame for the subtitle and navigation infos.
	frame := tview.NewFrame(tview.NewBox()).
		SetBorders(0, 0, 0, 0, 0, 0).
		AddText(subtitle, true, tview.AlignCenter, tcell.ColorDarkOrange).
		AddText("", true, tview.AlignCenter, tcell.ColorGold).
		AddText(navigation, true, tview.AlignCenter, tcell.ColorWhite)

	// Create a Flex layout that centers the logo and subtitle.
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tview.NewBox(), 0, 7, false).
		AddItem(tview.NewFlex().
			AddItem(tview.NewBox(), 0, 1, false).
			AddItem(logoBox, logoWidth, 1, true).
			AddItem(tview.NewBox(), 0, 1, false), logoHeight, 1, true).
		AddItem(frame, 0, 10, false)

	return flex
}
