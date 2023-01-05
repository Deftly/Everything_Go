package main

import (
  "fmt"
  "rsc.io/quote"
)

func main() {
  // Don't communicate by sharing memory, share memory by communicating.
  fmt.Println(quote.Go())
}
