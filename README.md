# qr

A lightweight QR encoder CLI and library.

* Simple to use
* PNG support
* Terminal support

## Install CLI

### Using GO
```shell
go install github.com/chey/qr

qr https://www.example.com
```

## Library usage
```shell
go get [-u] github.com/chey/qr
```

### Example

This example prints a small QR code to the terminal using the lowest level of error correction.

```go
package main

import (
    "github.com/chey/qr/code"
)

func main() {
    qr, _ := code.New("https://www.example.com", code.Low)
    qr.Print()
}
```
