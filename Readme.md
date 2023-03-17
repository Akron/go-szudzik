# go-szudzik

[![GoDoc](https://godoc.org/github.com/Akron/go-szudzik?status.svg)](https://godoc.org/github.com/Akron/go-szudzik)

This is a Go implementation of the [Szudzik](https://szudzik.com/ElegantPairing.pdf) elegant pairing function. It is a function to encode two natural numbers into a single one and vice versa.

# Usage

```go
package main

import (
  "fmt"

  "github.com/Akron/go-szudzik"
)

func main () {
  sz := szudzik.Pair(1, 2)
  fmt.Println(sz) // 3

  a, b := szudzik.Unpair(3)
  fmt.Println(a, b) // 1 2
}
```