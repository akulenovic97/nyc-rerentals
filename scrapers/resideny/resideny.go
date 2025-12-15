package resideny

import (
	"github.com/akulenovic97/nyc-rerentals/scrapers"
	"github.com/gocolly/colly/v2"
)

type ResideNYScraper struct {
	URL string
}

func New() *ResideNYScraper {
	return &ResideNYScraper{
		URL: "https://residenewyork.com/my-properties/?status=open-to-applications",
	}
}

func (s *ResideNYScraper) Scrape() ([]scrapers.Listing, error) {
	var listings []scrapers.Listing
	c := colly.NewCollector()

	c.OnHTML(".rh_page__listing article h3 a", func(e *colly.HTMLElement) {
		listing := scrapers.Listing{
			Name: e.Text,
			Link: e.Attr("href"),
		}
		listings = append(listings, listing)
	})

	err := c.Visit(s.URL)
	return listings, err
}