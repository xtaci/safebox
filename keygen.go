package main

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"os"
	"path/filepath"
)

func rawOutput(primitive tview.Primitive) *tview.Flex {
	return tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(nil, 0, 1, false).
			AddItem(primitive, 1, 1, true).
			AddItem(nil, 0, 1, false), 1, 1, true)
}

func showKeyGenPasswordPrompt(newkey *MasterKey, parent string, path string) {
	const (
		windowName   = "showKeyGenPasswordPrompt"
		windowWidth  = 40
		windowHeight = 10
		windowTitle  = "SET MASTERKEY PASSWORD 🔑"
	)

	form := tview.NewForm()
	form.SetTitle(windowTitle)
	form.SetBorder(true)
	passwordField := tview.NewInputField().SetLabel("Password").
		SetFieldWidth(64).
		SetMaskCharacter('*')

	passwordFieldConfirm := tview.NewInputField().SetLabel("Password Confirm").
		SetFieldWidth(64).
		SetMaskCharacter('*')

	form.AddFormItem(passwordField)
	form.AddFormItem(passwordFieldConfirm)

	form.AddButton("OK", func() {
		if passwordField.GetText() != passwordFieldConfirm.GetText() {
			showFailWindow("PASSWORD MISMATCH")
		} else {
			newkey.changePassword([]byte(passwordField.GetText()))
			err := newkey.store(path)
			// display message after store
			if err != nil {
				showFailWindow(err.Error())
			} else {
				newkey.path = path
				showSuccessWindow(fmt.Sprint("Successfully Stored Master Key!!!\n", path), func() {
					// set masterkey to newkey and update view
					masterKey = newkey
					refresh()
					layoutRoot.RemovePage(windowName)
					layoutRoot.RemovePage(parent)
				})
			}
		}
	})
	form.SetFocus(0)

	layoutRoot.AddPage(windowName, popup(windowWidth, windowHeight, form), true, true)
}

func showKeyEntropyInputWindow() {
	const (
		windowName   = "showKeyEntropyInputWindow"
		windowWidth  = 100
		windowHeight = 12
		windowTitle  = "- HIT KEYBOARD RANDOMLY -"
	)
	keyboardHits := uint8(0)
	var entropy string
	entropyText := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter).
		SetWrap(false).SetText("0/16")
	entropyText.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		keyboardHits++
		key := event.Key()
		entropy = fmt.Sprintf("%s|%d", entropy, key)
		s := fmt.Sprintf("%d/16", keyboardHits)
		entropyText.SetText(s)
		if keyboardHits == 16 {
			newKey := newMasterKey()
			sum := md5.Sum([]byte(entropy))
			newKey.generateMasterKey(sum[:])
			layoutRoot.RemovePage(windowName)
			showKeyGenWindow(newKey)
		}
		return nil
	})

	form := tview.NewForm()
	form.SetFocus(0)

	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow).
		SetTitle(windowTitle).
		SetBorder(true)
	text := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignLeft).
		SetWrap(false)

	flex.AddItem(text, 0, 1, false)
	flex.AddItem(entropyText, 0, 1, true)
	flex.AddItem(form, 0, 1, false)

	layoutRoot.AddPage(windowName, popup(windowWidth, windowHeight, flex), true, true)
}

func showKeyGenWindow(newkey *MasterKey) {
	const (
		windowName   = "showKeyGenWindow"
		windowWidth  = 100
		windowHeight = 12
		windowTitle  = "- MASTERKEY GENERATION -"
	)

	text := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignLeft).
		SetWrap(false)

	fmt.Fprint(text, "GENERATED MASTER KEY SHA256:\n\n")
	md := sha256.Sum256(newkey.masterKey[:])
	fmt.Fprintf(text, "[darkorange::]%v[-:-:-]\n", hexutil.Encode(md[:]))
	fmt.Fprint(text, "[white::l]MAKE SURE YOU BACKUP THIS FILE CORRECTLY\n")

	// path input field
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	form := tview.NewForm()
	inputField := tview.NewInputField().
		SetLabel("Save To:").
		SetText(filepath.Join(path, ".safebox.key")).
		SetFieldWidth(64)
	form.AddFormItem(inputField)
	form.AddButton("Save", func() {
		// check file existence
		if _, err := os.Stat(inputField.GetText()); os.IsNotExist(err) {
			showKeyGenPasswordPrompt(newkey, windowName, inputField.GetText())
		} else {
			showFailWindow("MASTER KEY FILE EXISTS, IF YOU WANT TO OVERWRITE, PLEASE DELETE THIS FILE BY YOURSELF.")
		}
	})
	form.AddButton("...", func() {
		showDirWindow(inputField)
	})

	form.SetFocus(2)

	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow).
		SetTitle(windowTitle).
		SetBorder(true)
	flex.AddItem(text, 0, 1, false)
	flex.AddItem(form, 0, 1, true)

	layoutRoot.AddPage(windowName, popup(windowWidth, windowHeight, flex), true, true)
}
