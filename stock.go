package gosupreme

import (
	"net/http"
)

type Stock struct {
	Categories          Categories `json:"products_and_categories"`
	ReleaseDate         string     `json:"release_date"`
	ReleaseWeek         string     `json:"release_week"`
	LastMobileAPIUpdate string     `json:"last_mobile_api_update"`
}

// Fetches the latest stop from supreme.
func (s *Supreme) FetchStock() error {
	s.l.Println("Fetching latest stock from supreme.")

	// create a new http GET request
	req, err := http.NewRequest("GET", STOCK_URL, nil)
	if err != nil {
		return err
	}

	// make the http request and save its response
	var stock *Stock
	err = s.makeRequest(req, &stock)
	if err != nil {
		return err
	}

	// update stock
	s.Stock = stock
	s.l.Printf("Fetched stock with release date '%s'.\n", s.Stock.ReleaseDate)
	return nil
}
