package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/rivo/tview"
)

var verString = "VER 1.0"
var mainFrameTitle = fmt.Sprintf("- SAFEBOX KEY MANGEMENT SYSTEM %v -", verString)

func exporterSelect(idx uint16) *tview.Flex {
	exporterLabel := "Select an exporter (hit Enter):"
	var exportorNames []string
	for k := range exports {
		exportorNames = append(exportorNames, exports[k].Name())
	}

	outputBox := tview.NewTextView()
	outputBox.SetScrollable(true)
	outputBox.SetWrap(true)

	selected := 0
	form := tview.NewForm()
	form.SetTitle("EXPORT KEY")
	form.AddDropDown(exporterLabel, exportorNames, 0, func(option string, optionIndex int) {
		selected = optionIndex
	})
	form.AddButton("Export", func() {
		key, _ := masterKey.deriveKey(idx, exports[selected].KeySize())
		bts, _ := exports[selected].Export(key)
		outputBox.Write(bts)
		form.RemoveButton(0)
		form.AddButton("Return", func() {
			root.RemovePage("export")
		})
	})
	form.SetFocus(0)

	view := tview.NewFlex()
	view.SetBorder(true)
	view.SetDirection(tview.FlexRow).
		AddItem(form, 0, 1, true).
		AddItem(outputBox, 0, 1, false)

	return modal(80, 20, view)
}

func deriveKeyOperation(idx uint16) *tview.Flex {
	form := tview.NewForm()
	form.SetBorder(true)
	form.SetTitle("SETTING KEY PROPERTIES")
	form.AddInputField("Label", masterKey.labels[idx], 16, nil, nil)
	form.AddButton("Update", func() {
		root.RemovePage("prompt")
		root.SwitchToPage("main")
	})
	form.AddButton("Export", func() {
		root.AddAndSwitchToPage("export", exporterSelect(idx), true)
	})
	form.SetFocus(0)

	return modal(40, 10, form)
}

// main operation frame
func mainFrameWindow() (content *tview.Flex) {
	if masterKey == nil {
		// if master key has not loaded
		text := tview.NewTextView()
		text.SetDynamicColors(true).
			SetTextAlign(tview.AlignCenter)
		fmt.Fprintf(text, `[red]KEY NOT LOADED
PLEASE LOAD A MASTER KEY[yellow][F2][red] OR GENERATE ONE[yellow][F1][red] FIRST`)

		flex := tview.NewFlex()
		flex.SetDirection(tview.FlexRow).
			SetBorder(true).
			SetTitle(mainFrameTitle)

		flex.AddItem(tview.NewBox(), 0, 8, false)
		flex.AddItem(text, 0, 1, true)
		flex.AddItem(tview.NewBox(), 0, 8, false)
		return flex
	}

	// key list
	list := tview.NewList()
	for i := uint16(0); i < 100; i++ {
		// we derive and show the part of the key
		key, err := masterKey.deriveKey(i, 32)
		if err != nil {
			log.Fatal(err)
		}

		var label string
		if masterKey.labels[i] != "" {
			label = masterKey.labels[i]
		} else {
			label = hex.EncodeToString(key[:8])
		}
		list.AddItem(fmt.Sprintf("KEY%v", i), label, 0, nil)
	}

	// key selection
	list.SetSelectedFunc(func(idx int, mainText, secondaryText string, shortcut rune) {
		root.AddAndSwitchToPage("prompt", deriveKeyOperation(uint16(idx)), true)
	})

	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow).
		SetBorder(true).
		SetTitle(mainFrameTitle)

	flex.AddItem(list, 0, 1, true)
	return flex
}
