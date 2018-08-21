package main

import (
  "errors"
  "log"
)

func main() {
  var errFatal = errors.New("this is a error")
  log.Fatal(errFatal)
}
