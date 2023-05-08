package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	os.Args = []string{"qr", "chey"}
	main()
}

func TestMainDefaults(t *testing.T) {
	var buf bytes.Buffer
	os.Args = []string{"qr", "chey"}
	flag.Parse()
	if err := execute(&buf); err != nil {
		t.Error(err)
	}
}

func TestMainBig(t *testing.T) {
	var buf bytes.Buffer
	os.Args = []string{"qr", "-big", "chey"}
	flag.Parse()
	if err := execute(&buf); err != nil {
		t.Error(err)
	}
}

func TestMainLevel(t *testing.T) {
	var buf bytes.Buffer
	os.Args = []string{"qr", "-level", "3", "chey"}
	flag.Parse()
	if err := execute(&buf); err != nil {
		t.Error(err)
	}
}

func TestMainErrors(t *testing.T) {
	var buf bytes.Buffer
	cases := []struct {
		args     []string
		expected error
	}{
		{[]string{"qr"}, ErrMissingText},
		{[]string{"qr", "chey", "wuz", "here"}, ErrToManyArgs},
		{[]string{"qr", "-level", "5", "chey"}, ErrInvalidLevel},
	}

	for _, c := range cases {
		os.Args = c.args
		flag.Parse()
		if err := execute(&buf); err == nil {
			t.Errorf("expected = %#v", c.expected)
		}
	}
}

func TestMainPng(t *testing.T) {
	var buf bytes.Buffer
	file := fmt.Sprintf("%s%c%s", t.TempDir(), os.PathSeparator, "qrcode.png")
	os.Args = []string{"qr", "-png", "-file", file, "-level", "0", "chey"}
	flag.Parse()
	if err := execute(&buf); err != nil {
		t.Error(err)
	}
}

func TestMainPngStdout(t *testing.T) {
	var buf bytes.Buffer
	os.Args = []string{"qr", "-png", "-file", "-", "-level", "0", "chey"}
	flag.Parse()

	rc := capture()
	if err := execute(&buf); err != nil {
		t.Error(err)
	}
	out := <-rc()

	if out == "" {
		t.Error("cannot send PNG to stdout")
	}
}
func TestMainHelp(t *testing.T) {
	fs := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	fs.Usage = usage
	err := fs.Parse([]string{"-h"})
	if err != flag.ErrHelp {
		t.Error("expected = flag.ErrHelp")
	}
}

func TestMainVersionFlag(t *testing.T) {
	var buf bytes.Buffer
	os.Args = []string{"qr", "-version"}
	flag.Parse()
	if err := execute(&buf); err != nil {
		t.Error(err)
	}
}

func capture() func() <-chan string {
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	c := make(chan string)

	// this copies stdout to the buffer without blocking
	go func() {
		var buf bytes.Buffer
		if _, err := io.Copy(&buf, r); err != nil {
			panic(err)
		}
		c <- buf.String()
	}()
	return func() <-chan string {
		defer func() {
			w.Close()
			os.Stdout = stdout
		}()
		return c
	}
}
