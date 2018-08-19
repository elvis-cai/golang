package main

import (
  "fmt"
  "log"
  "net/http"
  "time"
)

func responseTime(url string) {
  start := time.Now()
  res, err := http.Get(url)
  if err != nil {
    log.Fatal(err)
  }
  defer res.Body.Close()

  elapsed := time.Since(start).Seconds()
  fmt.Printf("%s tool %v seconds \n", url, elapsed)

}

func main() {
  urls := make([]string, 3)
  urls[0] = "https://www.usa.gov"
  urls[1] = "https://www.gov.uk"
  urls[2] = "https://www.youtube.com"

  for _, u := range urls {
  responseTime(u)
  }
}
