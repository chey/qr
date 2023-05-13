package code

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"os"

	"rsc.io/qr"
)

const (
	emptyBlock     = "\u0020"
	upperHalfBlock = "\u2580"
	lowerHalfBlock = "\u2584"
	fullBlock      = "\u2588"
	termBlack      = "\033[40m  \033[0m"
	termWhite      = "\033[47m  \033[0m"
	newline        = "\n"
)

type Level = qr.Level

const (
	Low     = qr.L
	Medium  = qr.M
	High    = qr.Q
	Highest = qr.H
)

// Code wraps module import
type Code struct {
	*qr.Code
}

// New creates a new QR Code.
// Error correction levels are 0, 1, 2, 3.
// 3 being the highest level of correction.
func New(text string, level Level) (*Code, error) {
	c, err := qr.Encode(text, level)
	if err != nil {
		return nil, err
	}
	return &Code{c}, nil
}

// Small writes Code to io.Writer w.
// Perfect for terminal.
func (c *Code) Small(w io.Writer) {
	for y := -1; y <= c.Size; y += 2 {
		for x := -1; x <= c.Size; x++ {
			next := c.Black(x, y+1)
			curr := c.Black(x, y)
			if curr && next {
				writeString(w, emptyBlock)
			} else if curr && !next {
				writeString(w, lowerHalfBlock)
			} else if !curr && !next {
				if y == c.Size {
					writeString(w, upperHalfBlock)
				} else {
					writeString(w, fullBlock)
				}
			} else {
				writeString(w, upperHalfBlock)
			}
		}
		if y < c.Size {
			writeString(w, newline)
		}
	}
}

// Big writes a larger Code to io.Writer w.
// Great for terminal.
func (c *Code) Big(w io.Writer) {
	for y := -1; y <= c.Size; y++ {
		for x := -1; x <= c.Size; x++ {
			if c.Black(x, y) {
				writeString(w, termBlack)
			} else {
				writeString(w, termWhite)
			}
		}
		if y < c.Size {
			writeString(w, newline)
		}
	}
}

// Print writes Code to os.Stdout
func (c *Code) Print() {
	c.Small(os.Stdout)
	writeString(os.Stdout, newline)
}

// Image returns an image.Image
func (c *Code) Image() image.Image {
	return &codeImage{c.Code}
}

// PNG writes a PNG image to io.Writer w
func (c *Code) PNG(w io.Writer) error {
	return png.Encode(w, c.Image())
}

// writeString writes string s to io.Writer w
func writeString(w io.Writer, s string) {
	if _, err := io.WriteString(w, s); err != nil {
		panic(err)
	}
}

// codeImage implements image.Image
type codeImage struct {
	*qr.Code
}

var (
	whiteColor color.Color = color.Gray{0xFF}
	blackColor color.Color = color.Gray{0x00}
)

func (c *codeImage) Bounds() image.Rectangle {
	d := (c.Size + 8) * c.Scale
	return image.Rect(0, 0, d, d)
}

func (c *codeImage) At(x, y int) color.Color {
	if c.Black(x/c.Scale-4, y/c.Scale-4) {
		return blackColor
	}
	return whiteColor
}

func (c *codeImage) ColorModel() color.Model {
	return color.GrayModel
}
