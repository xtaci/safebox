package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
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
		windowHeight = 3
		windowTitle  = "- TYPE IN RANDOM KEYS TO GATHER ENTROPY -"
	)

	// a sinker
	chanClosed := make(chan struct{})
	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow).
		SetTitle(windowTitle).
		SetBorder(true)

	var keyboardHits int32
	var entropy string

	entropyText := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignLeft).
		SetWrap(false).SetText("[0%]")

	entropyText.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// display hits
		hits := atomic.AddInt32(&keyboardHits, 1)
		_, _, width, _ := entropyText.GetInnerRect()
		s := fmt.Sprintf("[blue][%s %d%%]", strings.Repeat("#", int(float32(width-7)*float32(hits)/100)), hits)
		entropyText.SetText(s)

		// gather entropy
		key := event.Key()
		entropy = fmt.Sprintf("%s|%d|%d", entropy, key, time.Now().UnixNano())

		// entropy enough
		if hits >= 100 {
			close(chanClosed)
			showSuccessWindow(fmt.Sprint("Successfully Generated Master Key"), func() {
				newKey := newMasterKey()
				newKey.generateMasterKey([]byte(entropy))
				layoutRoot.RemovePage(windowName)
				showKeySaveWindow(newKey)
			})
		}
		return nil
	})

	text := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignLeft).
		SetWrap(false)

	flex.AddItem(text, 0, 1, false)
	flex.AddItem(entropyText, 0, 1, true)

	go func() {
		ticker := time.NewTicker(time.Second)
		for {
			select {
			case <-ticker.C:
				hits := atomic.AddInt32(&keyboardHits, -1)
				if hits < 0 {
					hits = atomic.AddInt32(&keyboardHits, 1)
				}
				app.QueueUpdateDraw(func() {
					_, _, width, _ := entropyText.GetInnerRect()
					s := fmt.Sprintf("[gray][%s %d%%]", strings.Repeat("#", int(float32(width-7)*float32(hits)/100)), hits)
					entropyText.SetText(s)
				})
			case <-chanClosed:
				return
			}
		}
	}()

	layoutRoot.AddPage(windowName, popup(windowWidth, windowHeight, flex), true, true)
}

func showKeySaveWindow(newkey *MasterKey) {
	const (
		windowName   = "showKeySaveWindow"
		windowWidth  = 100
		windowHeight = 12
		windowTitle  = "- MASTERKEY GENERATED -"
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
