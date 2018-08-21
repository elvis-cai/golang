package main

import (
  "log"
  "os"
)

func main() {
  f, err := os.OpenFile("haha.txt",
os.O_APPEND|os.O_CREATE|os.O_RDWR,0666)
  if err != nil {
    log.Fatal(err)
  }
  defer f.Close()
  log.SetOutput(f)
  for i:=1;i<=5;i++ {
    log.Printf("Log iteration %d", i)
  }
}
