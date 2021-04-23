package qrcode

import (
	"fmt"
	"testing"
)

func TestQRCode(t *testing.T) {
	fmt.Print(GenerateQRCode("hello world"))
}
