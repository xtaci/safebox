// +build !windows

package main

import (
	"os"
	"os/exec"
	"strings"
)

var (
	wideCharset = []string{"zh_", "jp_", "ko_", "ja_", "th_", "hi_"}
)

func fixCharset() {
	locale := os.Getenv("LANG")

	var asianCharset bool
	for k := range wideCharset {
		if strings.HasPrefix(locale, wideCharset[k]) {
			asianCharset = true
		}
	}

	if asianCharset {
		os.Setenv("LANG", "C.UTF-8")
		cmd := exec.Command(os.Args[0])
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
		os.Exit(0)
	}
}
