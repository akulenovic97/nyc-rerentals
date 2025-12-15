package main

import (
	"fmt"
	"log"
	"github.com/gocolly/colly/v2"
)

func main() {
	fmt.Println("NYC re-rentals scraper")

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Now visiting ", r.URL)
	})
	c.OnHTML(".rh_page__listing article h3 a", func(e *colly.HTMLElement) {
		name := e.Text
		link := e.Attr("href")
		fmt.Println("Property: ", name)
		fmt.Println("Link: ", link)
	})
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request failed: ", err)
	})
	c.Visit("https://residenewyork.com/my-properties/?status=open-to-applications")
}