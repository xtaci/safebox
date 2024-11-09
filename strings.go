package main

const (
	S_WINDOW_SHOWLOADPASSWORD_TITLE = " 🔓 Decrypting masterkey... "
	S_WINDOW_SHOW_LOADKEY_TITLE     = " Loading master key... "

	S_MAIN_FRAME_TITLE             = " SAFEBOX KEY MANAGEMENT SYSTEM "
	S_MAIN_FRAME_CELL_ID           = "Id"
	S_MAIN_FRAME_CELL_LABEL        = "Label"
	S_MAIN_FRAME_CELL_DERIVED_KEYS = "Derived Keys"

	S_WINDOW_EXPORT_TITLE  = " Export Key "
	S_WINDOW_EXPORT_LABEL  = "Select an exporter (hit Enter) -> "
	S_WINDOW_EXPORT_BUTTON = "Show"

	S_WINDOW_SETLABEL_TITLE  = " Setting key label "
	S_WINDOW_SETLABEL_BUTTON = "Update"

	S_WINDOW_CHANGEPASS_TITLE          = " 🔑 Changing masterkey password... "
	S_WINDOW_CHANGEPASS_LABEL_PASSWORD = "Password"
	S_WINDOW_CHANGEPASS_LABEL_CONFIRM  = "Confirm Password"
	S_WINDOW_CHANGEPASS_BUTTON_OK      = "OK"

	S_WINDOW_KEYGEN_PASSWORD_TITLE          = " 🔑 Setting masterkey password... "
	S_WINDOW_KEYGEN_PASSWORD_LABEL_PASSWORD = "Password"
	S_WINDOW_KEYGEN_PASSWORD_LABEL_CONFIRM  = "Confirm Password"
	S_WINDOW_KEYGEN_PASSWORD_BUTTON_OK      = "OK"

	S_WINDOW_ENTROPY_TITLE = " Please type in random keys for randomness... "

	S_WINDOW_KEYSAVE_TITLE        = " Saving generated masterkey... "
	S_WINDOW_KEYSAVE_LABEL_SAVETO = "Save to: "
	S_WINDOW_KEYSAVE_BUTTON_SAVE  = "Save"
	S_WINDOW_KEYSAVE_BUTTON_3DOTS = "..."

	S_WINDOW_SHOWDIR_TITLE = " Select a path to save masterkey... "

	S_MODAL_BUTTON_OK     = "OK"
	S_MODAL_TITLE_ERROR   = "<ERROR>"
	S_MODAL_TITLE_SUCCESS = "<SUCCESS>"
	S_MODAL_TITLE_INFO    = "<INFO>"

	S_MSG_PASSWORD_MISMATCH = "Password mismatch"
	S_MSG_PASSWORD_CHANGED  = "Masterkey Password changed"
)

const (
	S_INFOBOX_KEYINFO = `
[-:-:-]Version
[darkblue]%v

[-:-:-]Location:
[darkblue]%v

[-:-:-]Master Key SHA256:
[darkblue]%v

[-:-:-]Master Keys Created At:
[darkblue]%v

[-:-:-]Number of keys with label:
[darkblue]%v

[-:-:-]System:
[darkblue]%v %v

`

	S_INFOBOX_INSTRUCTIONS = `
Instructions

1) Use ArrowKeys [darkred]← ↑ → ↓ [-:-:-]To Select Keys, masks on derived keys will be uncovered when selected.

2) Press [darkred]<Enter>[-:-:-] on 'Derived Keys' column to export.

3) Press [darkred]<Enter>[-:-:-] on 'Label' column to set label.

4) Use [darkred]<TAB>[-:-:-] to focus on different items.
`
)

const (
	S_FOOTER_COPYRIGHT = "//safebox // Copyright (c) 2021 xtaci"
)
