package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var verString = "VER 1.0"
var mainFrameTitle = fmt.Sprintf("- SAFEBOX KEY MANGEMENT SYSTEM %v -", verString)

func showExporterSelect(idx uint16) {
	windowName := "showExporterSelect"
	exporterLabel := "Select an exporter (hit Enter):"
	var exportorNames []string
	for k := range exports {
		exportorNames = append(exportorNames, exports[k].Name())
	}

	selected := 0
	form := tview.NewForm()
	form.SetTitle("EXPORT KEY")
	form.AddDropDown(exporterLabel, exportorNames, 0, func(option string, optionIndex int) {
		selected = optionIndex
	})
	form.AddButton("Export", func() {
		key, _ := masterKey.deriveKey(idx, exports[selected].KeySize())
		bts, _ := exports[selected].Export(key)

		// output page
		outputBox := tview.NewTextView()
		outputBox.SetScrollable(true)
		outputBox.SetWrap(true)
		outputBox.Write(bts)
		root.AddPage("output", outputBox, true, true)
	})
	form.SetFocus(0)

	view := tview.NewFlex()
	view.SetBorder(true)
	view.SetDirection(tview.FlexRow).
		AddItem(form, 0, 1, true)

	root.AddPage(windowName, popup(100, 20, view), true, true)
}

func showDeriveKeyOperation(idx uint16) {
	windowName := "showDeriveKeyOperation"
	form := tview.NewForm()
	form.SetBorder(true)
	form.SetTitle("SETTING KEY PROPERTIES")
	form.AddInputField("Label", masterKey.lables[idx], 16, nil, nil)
	form.AddButton("Update", func() {
		//update key
		masterKey.lables[idx] = form.GetFormItemByLabel("Label").(*tview.InputField).GetText()
		masterKey.store(masterKey.password, masterKey.path)
		// refresh main frame
		mainFrame = mainFrameWindow()
		refreshBody()
		root.RemovePage(windowName)
	})
	form.AddButton("Export", func() {
		showExporterSelect(idx)
	})
	form.SetFocus(0)

	root.AddPage(windowName, popup(40, 10, form), true, true)
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

	// key table
	table := tview.NewTable().SetBorders(true)
	table.SetTitle(mainFrameTitle)

	// table header
	table.SetCell(0, 0,
		tview.NewTableCell("ID").
			SetTextColor(tcell.ColorRed).
			SetSelectable(false).
			SetAlign(tview.AlignLeft))

	table.SetCell(0, 1,
		tview.NewTableCell("KEY NAME").
			SetTextColor(tcell.ColorRed).
			SetSelectable(false).
			SetAlign(tview.AlignLeft))

	table.SetCell(0, 2,
		tview.NewTableCell("KEY PREFIX").
			SetTextColor(tcell.ColorRed).
			SetSelectable(false).
			SetAlign(tview.AlignLeft))

	// fix table header & first column
	table.SetFixed(1, 1)
	table.SetSelectable(true, true)

	addDerivedKeys := func(start uint16) {
		for i := uint16(0); i < 64; i++ {
			idx := start + i
			if idx >= MaxKeys {
				return
			}
			// we derive and show the part of the key
			key, err := masterKey.deriveKey(idx, 32)
			if err != nil {
				log.Fatal(err)
			}

			table.SetCell(int(idx)+1, 0,
				tview.NewTableCell(fmt.Sprint(idx)).
					SetTextColor(tcell.ColorWhite).
					SetAlign(tview.AlignLeft).
					SetSelectable(false))

			table.SetCell(int(idx)+1, 1,
				tview.NewTableCell(masterKey.lables[idx]).
					SetTextColor(tcell.ColorRed).
					SetAlign(tview.AlignLeft).
					SetSelectable(true))

			table.SetCell(int(idx)+1, 2,
				tview.NewTableCell(hex.EncodeToString(key)).
					SetTextColor(tcell.ColorRed).
					SetAlign(tview.AlignLeft))
		}
	}

	// add initial derived keys
	addDerivedKeys(0)

	table.SetSelectionChangedFunc(func(row, column int) {
		// moved to last
		if row == table.GetRowCount()-1 {
			addDerivedKeys(uint16(row) - 1)
		}
	})

	// key selection
	table.SetSelectedFunc(func(row, column int) {
		showDeriveKeyOperation(uint16(row) - 1)
	})

	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow).
		SetBorder(true).
		SetTitle(mainFrameTitle)

	flex.AddItem(table, 0, 1, true)
	return flex
}
