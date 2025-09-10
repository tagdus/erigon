package main

import (
  "fmt"
  erigonapp "github.com/erigontech/erigon/turbo/app"
)

func main() {
  fmt.Println("Hello, World from main()!")
  erigonapp.MakeApp()
}
