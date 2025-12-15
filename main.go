package main

import (
	"fmt"
	"log"
	"github.com/akulenovic97/nyc-rerentals/scrapers/resideny"
)

func main() {
	fmt.Println("NYC re-rentals scraper")

	scraper := resideny.New()
	listings, err := scraper.Scrape()

	if err != nil {
		log.Fatal("Error scraping: ", err)
	}

	for i, listing := range listings {
		fmt.Printf("%d. %s\n    %s\n\n", i + 1, listing.Name, listing.Link)
	}
	fmt.Printf("Total listings found: %d\n", len(listings))
}