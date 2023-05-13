package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/chey/qr/code"
)

var (
	ErrMissingText  = errors.New("missing text argument")
	ErrToManyArgs   = errors.New("too many arguments")
	ErrInvalidLevel = errors.New("invalid error correction level")
)

var (
	big   bool
	file  string
	level int
	png   bool
	ver   bool
)

func init() {
	flag.Usage = usage
	flag.BoolVar(&big, "big", false, "output big qr code")
	flag.IntVar(&level, "level", int(code.Medium), fmt.Sprintf("error correction level: %d, %d, %d, %d", code.Low, code.Medium, code.High, code.Highest))
	flag.StringVar(&file, "file", "qrcode.png", "output png to `filename`")
	flag.BoolVar(&png, "png", false, "output as png")
	flag.BoolVar(&ver, "version", false, "version for qr")
}

func main() {
	flag.Parse()
	if err := execute(os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}

func execute(w io.Writer) error {
	if ver {
		fmt.Println("qr", code.Version())
		return nil
	}

	if level < int(code.Low) || level > int(code.Highest) {
		return ErrInvalidLevel
	}

	if flag.NArg() == 0 {
		return ErrMissingText
	}

	if flag.NArg() > 1 {
		return ErrToManyArgs
	}

	qr, err := code.New(flag.Arg(0), code.Level(level))
	if err != nil {
		return err
	}

	if png && file == "-" {
		if err := qr.PNG(os.Stdout); err != nil {
			return err
		}
	} else if png {
		f, err := os.Create(file)
		if err != nil {
			return err
		}
		defer f.Close()
		if err := qr.PNG(f); err != nil {
			return err
		}
		fmt.Fprintf(os.Stderr, "png written to \"%s\"\n", file)
	} else if big {
		qr.Big(w)
		fmt.Fprintln(w)
	} else {
		qr.Small(w)
		fmt.Fprintln(w)
	}

	return nil
}

func usage() {
	fmt.Fprintf(flag.CommandLine.Output(), `qr %s, a lightweight QR encoder.
Usage: %s [FLAGS]... [TEXT]

Flags:
`, code.Version(), os.Args[0])
	flag.PrintDefaults()
	fmt.Fprintf(flag.CommandLine.Output(), `
Example:
  Display QR code in terminal window.

    %[1]s https://www.example.com

  Display PNG QR code on screen.

    %[1]s -png -file - https://www.example.com | display

`, os.Args[0])
}
