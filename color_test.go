package vt10x

import (
	"testing"
)

func TestByte2Color(t *testing.T) {
	color := byte2color(0)
	if color != 0x2e3436 {
		t.Fatal(color)
	}
	color = byte2color(15)
	if color != 0xeeeeec {
		t.Fatal(color)
	}
	color = byte2color(16)
	if color != 0x000000 {
		t.Fatal(color)
	}
	color = byte2color(231)
	if color != 0xffffff {
		t.Fatal(color)
	}
}
