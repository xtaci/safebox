package main

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestWinDir(t *testing.T) {
	dir := "C:\\Windows\\system32\\drivers"
	rootDir := "/"
	if runtime.GOOS == "windows" {
		rootDir = "C:\\"
	}

	for dir != rootDir {
		if dir == "." {
			wd, err := os.Getwd()
			if err != nil {
				panic(err)
			}
			dir = wd
		}

		dir = filepath.Dir(dir)
		t.Log(dir)
	}
}
