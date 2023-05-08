package main

/*
#include <stdlib.h>
typedef struct Png {
    int n;
	char *b;
} Png;
*/
import "C"

import (
	"bytes"

	"github.com/chey/qr/code"
)

//export qrCodeSmall
func qrCodeSmall(text *C.char, level int) *C.char {
	var buf bytes.Buffer
	c, _ := code.New(C.GoString(text), code.Level(level))
	c.Small(&buf)
	return C.CString(buf.String())
}

//export qrCodeBig
func qrCodeBig(text *C.char, level int) *C.char {
	var buf bytes.Buffer
	c, _ := code.New(C.GoString(text), code.Level(level))
	c.Big(&buf)
	return C.CString(buf.String())
}

//export qrCodePng
func qrCodePng(text *C.char, level int) C.Png {
	p := new(C.Png)
	var buf bytes.Buffer
	c, _ := code.New(C.GoString(text), code.Level(level))
	err := c.PNG(&buf)
	if err != nil {
		return *p
	}
	p.n = C.int(buf.Len())
	p.b = C.CString(buf.String())
	return *p
}

func main() {}
