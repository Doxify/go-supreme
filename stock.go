package gosupreme

import (
	"net/http"
	"reflect"
	"strings"
)

type Stock struct {
	Categories          Categories `json:"products_and_categories"`
	ReleaseDate         string     `json:"release_date"`
	ReleaseWeek         string     `json:"release_week"`
	LastMobileAPIUpdate string     `json:"last_mobile_api_update"`
}

func (s *Supreme) FetchStock() error {
	s.l.Println("Fetching latest stock from supreme.")

	// create a new http GET request
	req, err := http.NewRequest("GET", StockURL, nil)
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

func (s *Stock) getProductsByCategory(category string) *Products {
	fields := reflect.TypeOf(s.Categories)
	values := reflect.ValueOf(s.Categories)
	num := fields.NumField()

	// loop through all categories
	for i := 0; i < num; i++ {
		c := fields.Field(i)
		p := values.Field(i)

		// if category exists, return its products
		if strings.EqualFold(category, c.Name) {
			products, _ := p.Interface().(Products)
			return &products
		}
	}

	// return nil if category does not exist
	return nil
}

func (s *Stock) getProductsByKeyword(keyword string) *Products {
	var matchedProducts Products

	fields := reflect.TypeOf(s.Categories)
	values := reflect.ValueOf(s.Categories)
	num := fields.NumField()

	// loop through all categories
	for i := 0; i < num; i++ {
		p := values.Field(i)
		products, _ := p.Interface().(Products)

		// loop through all products in a category
		for _, product := range products {
			// find products which have the keyword in their names
			if strings.Contains(strings.ToLower(product.Name), strings.ToLower(keyword)) {
				matchedProducts = append(matchedProducts, product)
			}
		}

	}

	return &matchedProducts
}

func (s *Stock) getProductsByCategoryAndKeyword(keyword string, category string) *Products {
	// storing all products that match the keyword and category
	var matchedProducts Products

	// getting all products that match the category
	products := s.getProductsByCategory(category)
	if products == nil {
		return nil
	}

	// finding products in the category that match the keyword
	for _, p := range *products {
		if strings.Contains(strings.ToLower(p.Name), strings.ToLower(keyword)) {
			matchedProducts = append(matchedProducts, p)
		}
	}

	return &matchedProducts
}
