package qrcode

import (
	"fmt"
	"testing"
)

func TestQRCode(t *testing.T) {
	fmt.Print(generateQRCode("hello world"))
}
