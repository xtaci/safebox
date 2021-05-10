// +build windows

package main

import (
	"os"
	"os/exec"
)

const codePageSet = "SAFEBOX_CODEPAGE_SET"

func fixCharset() {
	if os.Getenv(codePageSet) != "1" {
		exec.Command("CHCP", "65001").Run()
		os.Setenv(codePageSet, "1")
		cmd := exec.Command(os.Args[0])
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
		os.Exit(0)
	}
}
