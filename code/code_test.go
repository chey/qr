package code

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestCode(t *testing.T) {
	_, err := New("chey", 0)
	if err != nil {
		t.Error(err)
	}
}

// max is 7089 for numeric characters so this should fail
// http://en.wikipedia.org/wiki/QR_code
func TestCodeError(t *testing.T) {
	if _, err := New(strings.Repeat("1", 7090), 0); err == nil {
		t.Error("expeced error")
	}
}

func TestCodeSmall(t *testing.T) {
	c, _ := New("chey", 0)

	var buf bytes.Buffer

	c.Small(&buf)

	if buf.String() == "" {
		t.Error("got empty code")
	}

	got := buf.Len()
	want := 707

	if buf.Len() != want {
		t.Errorf("Buffer.Len() == %d, want = %d", got, want)
	}
}

func TestCodeBig(t *testing.T) {
	c, _ := New("chey", 0)

	var buf bytes.Buffer

	c.Big(&buf)

	if buf.String() == "" {
		t.Error("got empty code")
	}

	got := buf.Len()
	want := 5841

	if buf.Len() != want {
		t.Errorf("Buffer.Len() == %d, want = %d", got, want)
	}
}

func TestCodePNG(t *testing.T) {
	c, _ := New("chey", 0)

	var buf bytes.Buffer

	err := c.PNG(&buf)
	if err != nil {
		t.Error(err)
	}

	if buf.String() == "" {
		t.Error("got empty png")
	}
}

func TestCodePrint(t *testing.T) {
	r, w, _ := os.Pipe()
	orig := os.Stdout
	os.Stdout = w
	buf := make([]byte, 1024)

	c, _ := New("chey", 0)

	c.Print()
	n, err := r.Read(buf)
	if err != nil {
		t.Error(err)
	}
	os.Stdout = orig

	want := 708
	if n != want {
		t.Errorf("Print(chey) = %q, want = %q", n, want)
	}
}

func TestCodeWriteStringPanic(t *testing.T) {
	file := fmt.Sprintf("%s%c%s", t.TempDir(), os.PathSeparator, "qrcode.txt")
	fmt.Println(file)
	f, _ := os.Create(file)
	if err := f.Chmod(0400); err != nil {
		t.Error(err)
	}
	f.Close()
	f, _ = os.Open(file)

	recovered := false

	defer func() {
		if r := recover(); r != nil {
			recovered = true
			fmt.Println("Recovered in TestCodeWriteStringPanic", r)
		}
	}()
	writeString(f, "hi")

	if !recovered {
		t.Error("expected we could handle a write panic but we didn't")
	}
}
