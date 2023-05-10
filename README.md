# qr

[![Docker](https://github.com/chey/qr/actions/workflows/docker-publish.yml/badge.svg)](https://github.com/chey/qr/actions/workflows/docker-publish.yml)
[![goreleaser](https://github.com/chey/qr/actions/workflows/goreleaser.yml/badge.svg)](https://github.com/chey/qr/actions/workflows/goreleaser.yml)

A lightweight QR encoder CLI and library.

* Simple to use
* PNG support
* Terminal support

## Install and Use CLI

### Using GO
```shell
go install github.com/chey/qr@latest

qr https://www.example.com
```
![image](https://user-images.githubusercontent.com/152618/236944628-d9d0b7d2-14f7-4f40-b1ee-fd5640b6264a.png)

### Docker
```
docker run --rm ghcr.io/chey/qr https://www.example.com
```

### Brew
```
brew install chey/tap/qr
```

### Other
See the [releases](https://github.com/chey/qr/releases) page for other downloads.

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
