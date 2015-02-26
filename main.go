package main

import (
  "fmt"
  "log"

  "github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
  links := make(map[string]int)

  doc, err := goquery.NewDocument("http://news.ycombinator.com")
  if err != nil {
    log.Fatal(err)
  }

  links["http://news.ycombinator.com"] = 0
  doc.Find("td.title a").Each(func(i int, s *goquery.Selection) {
    // s.Text()
    links["http://news.ycombinator.com"]++
    //fmt.Printf("Link: %s\n", name)
  })

  fmt.Printf("%d\n", links["http://news.ycombinator.com"])
}

func main() {
  ExampleScrape()
}
