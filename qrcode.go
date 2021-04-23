package main

import (
	"bytes"
	"fmt"

	"github.com/skip2/go-qrcode"
)

func generateQRCode(text string) string {
	qr, err := qrcode.New(text, qrcode.Medium)
	if err != nil {
		panic(err)
	}

	bitmap := qr.Bitmap()
	row := len(bitmap)
	column := len(bitmap[0])

	var out bytes.Buffer
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if bitmap[i][j] {
				fmt.Fprint(&out, "\u2588\u2588")
			} else {
				fmt.Fprint(&out, "  ")
			}
		}
		fmt.Fprintln(&out, "")
	}

	return out.String()
}
