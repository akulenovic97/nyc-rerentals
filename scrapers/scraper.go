package scrapers

type Listing struct {
	Name string
	Link string
}

type Scraper interface {
	Scrape() ([]Listing, error)
}