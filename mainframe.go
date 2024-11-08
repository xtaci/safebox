package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func showExporterWindow(row int, col int) {
	windowName := "showExporterWindow"
	var exportorNames []string
	for k := range exports {
		exportorNames = append(exportorNames, exports[k].Name())
	}

	desc := tview.NewTextView()
	desc.SetWordWrap(true)
	desc.SetDynamicColors(true)
	desc.SetWrap(true)

	idx := uint16(row - 1)
	selected := 0
	form := tview.NewForm()
	form.SetTitle(S_WINDOW_EXPORT_TITLE)

	dropdown := tview.NewDropDown()
	dropdown.SetLabel(S_WINDOW_EXPORT_LABEL)
	dropdown.SetOptions(exportorNames, func(option string, optionIndex int) {
		selected = optionIndex
		desc.Clear()
		fmt.Fprintf(desc, exports[selected].Desc())
	})

	dropdown.SetCurrentOption(0)
	form.AddFormItem(dropdown)

	form.AddButton(S_WINDOW_EXPORT_BUTTON, func() {
		key, _ := masterKey.deriveKey(idx, exports[selected].KeySize())
		bts, _ := exports[selected].Export(key)

		// output page
		outputBox := tview.NewTextView()
		outputBox.SetDynamicColors(true)
		outputBox.SetScrollable(true)
		outputBox.SetWrap(true)
		outputBox.Write(bts)
		layoutRoot.AddPage("output", outputBox, true, true)
	})
	form.SetFocus(0)

	view := tview.NewFlex()
	view.SetBorder(true)
	view.SetDirection(tview.FlexRow).
		AddItem(form, 0, 1, true).
		AddItem(desc, 0, 2, true)

	layoutRoot.AddPage(windowName, popup(80, 20, view), true, true)
}

func showSetLabelWindow(row int, col int) {
	windowName := "showSetLabelWindow"
	idx := uint16(row - 1)

	form := tview.NewForm()
	form.SetBorder(true)
	form.SetTitle(S_WINDOW_SETLABEL_TITLE)
	form.AddInputField("Label", masterKey.labels[idx], 16, nil, nil)
	form.AddButton(S_WINDOW_SETLABEL_BUTTON, func() {
		//update key
		masterKey.setLabel(idx, form.GetFormItemByLabel("Label").(*tview.InputField).GetText())
		masterKey.store(masterKey.path)
		layoutMainBodyTable.GetCell(int(idx)+1, 1).Text = masterKey.labels[idx]
		layoutRoot.RemovePage(windowName)
		refreshInfo()
	})
	form.SetFocus(0)

	layoutRoot.AddPage(windowName, popup(40, 10, form), true, true)
}

// main operation frame
func mainFrameWindow() (content *tview.Flex) {
	layoutMainBody = tview.NewFlex()
	layoutMainBody.SetDirection(tview.FlexRow).SetBorder(true)
	refreshMainFrame()
	return layoutMainBody
}

func refreshMainFrame() {
	layoutMainBody.Clear()

	if masterKey == nil {
		// if master key has not loaded
		text := tview.NewTextView()
		text.SetDynamicColors(true).
			SetTextAlign(tview.AlignCenter)
		fmt.Fprintf(text, `[red]Master Key not loaded,
Please Load a master key or Generate one first
`)
		lang := os.Getenv("LANG")
		if strings.HasPrefix(lang, "zh_") {
			fmt.Fprintf(text, `[blue::bl]
Tips:
To get the BEST display quality, please export LANG to non Asian Language with UTF-8, such as:
export LANG=C.UTF-8
export LANG=en_US.UTF-8
`)
		}

		layoutMainBody.AddItem(tview.NewBox(), 0, 4, false)
		layoutMainBody.AddItem(text, 0, 2, true)
		layoutMainBody.AddItem(tview.NewBox(), 0, 4, false)
		return
	}

	// disable border, show table only
	layoutMainBody.SetBorder(false)

	// key table
	layoutMainBodyTable = tview.NewTable().SetBorders(true).SetBordersColor(tcell.ColorWhite)

	// table header
	layoutMainBodyTable.SetCell(0, 0,
		tview.NewTableCell(S_MAIN_FRAME_CELL_ID).
			SetTextColor(tcell.ColorWhite).
			SetBackgroundColor(tcell.ColorDarkGreen).
			SetSelectable(false).
			SetAlign(tview.AlignLeft))

	layoutMainBodyTable.SetCell(0, 1,
		tview.NewTableCell(S_MAIN_FRAME_CELL_LABEL).
			SetTextColor(tcell.ColorWhite).
			SetBackgroundColor(tcell.ColorDarkGreen).
			SetSelectable(false).
			SetAlign(tview.AlignLeft))

	layoutMainBodyTable.SetCell(0, 2,
		tview.NewTableCell(S_MAIN_FRAME_CELL_DERIVED_KEYS).
			SetTextColor(tcell.ColorWhite).
			SetBackgroundColor(tcell.ColorDarkGreen).
			SetSelectable(false).
			SetAlign(tview.AlignLeft))

	// fix table header & first column
	layoutMainBodyTable.SetFixed(1, 1)
	layoutMainBodyTable.SetSelectable(true, true)

	addDerivedKeys := func(start uint16) {
		for i := uint16(0); i < 64; i++ {
			idx := start + i
			// we derive and show the part of the key
			key, err := masterKey.deriveKey(idx, 32)
			if err != nil {
				panic(err)
			}

			layoutMainBodyTable.SetCell(int(idx)+1, 0,
				tview.NewTableCell(fmt.Sprint(idx)).
					SetAlign(tview.AlignLeft).
					SetTextColor(tcell.ColorWhite).
					SetBackgroundColor(tcell.ColorDarkGreen).
					SetSelectable(false))

			layoutMainBodyTable.SetCell(int(idx)+1, 1,
				tview.NewTableCell(masterKey.labels[idx]).
					SetTextColor(tcell.ColorWhite).
					SetBackgroundColor(tcell.ColorDarkGreen).
					SetAlign(tview.AlignLeft).
					SetSelectable(true))

			layoutMainBodyTable.SetCell(int(idx)+1, 2,
				tview.NewTableCell(mask(hex.EncodeToString(key), 4, '*')).
					SetTextColor(tcell.ColorLightGreen).
					SetBackgroundColor(tcell.ColorDarkGreen).
					SetAlign(tview.AlignLeft))
		}
	}

	// add initial derived keys
	addDerivedKeys(0)

	var lastRow int
	var lastCol int
	layoutMainBodyTable.SetSelectionChangedFunc(func(row, column int) {
		// moved to last
		idx := uint16(row) - 1
		if row == layoutMainBodyTable.GetRowCount()-1 {
			addDerivedKeys(idx)
		}

		// mask previous key again
		// derive key again
		if lastRow > 0 {
			key, err := masterKey.deriveKey(uint16(lastRow)-1, 32)
			if err != nil {
				panic(err)
			}

			layoutMainBodyTable.SetCell(lastRow, lastCol,
				tview.NewTableCell(mask(hex.EncodeToString(key), 4, '*')).
					SetTextColor(tcell.ColorLightGreen).
					SetBackgroundColor(tcell.ColorDarkGreen).
					SetAlign(tview.AlignLeft))
		}

		// mask previous selection
		if column == 2 && row > 0 {
			// derive key again
			key, err := masterKey.deriveKey(idx, 32)
			if err != nil {
				panic(err)
			}

			// uncover mask
			layoutMainBodyTable.SetCell(row, column,
				tview.NewTableCell(hex.EncodeToString(key)).
					SetTextColor(tcell.ColorLightGreen).
					SetBackgroundColor(tcell.ColorDarkGreen).
					SetAlign(tview.AlignLeft))
			// remember last selection
			lastRow = row
			lastCol = column
		}
	})

	// key selection
	layoutMainBodyTable.SetSelectedFunc(func(row, column int) {
		if column == 1 {
			showSetLabelWindow(row, column)
		} else if column == 2 {
			showExporterWindow(row, column)
		}
	})

	layoutMainBody.AddItem(layoutMainBodyTable, 0, 1, true)
}
