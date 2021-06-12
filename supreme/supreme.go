package supreme

import (
	"net/http"
)

type Supreme struct {
	Client *Client
	Stock  *Stock
}

func CreateNewInstance() *Supreme {
	// create a new client
	c := NewClient("supremego")
	c.Logger.Println("Successfully created SupremeGo instance.")
	return &Supreme{
		Client: c,
		Stock:  nil,
	}
}

func (s *Supreme) FetchStock() error {
	s.Client.Logger.Println("Fetching latest stock from supreme.")

	// create a new http GET request
	req, err := http.NewRequest("GET", s.Client.BaseURL, nil)
	if err != nil {
		return err
	}

	// make the http request and save its response
	var stock *Stock
	err = s.Client.makeRequest(req, &stock)
	if err != nil {
		return err
	}

	// update stock
	s.Stock = stock
	s.Client.Logger.Printf("Fetched stock with release date '%s'.\n", s.Stock.ReleaseDate)
	return nil
}

func (s *Supreme) GetAllCategoryProducts(category string) *Products {
	s.Client.Logger.Printf("Fetching all products from the '%s' category...", category)
	return s.Stock.getProductsByCategory(category)
}

func (s *Supreme) GetAllProducts(keyword string) *Products {
	s.Client.Logger.Printf("Fetching all products with the '%s' keyword...", keyword)
	return s.Stock.getProductsByKeyword(keyword)
}

func (s *Supreme) LookupCategoryProduct(keyword string, category string) *Products {
	s.Client.Logger.Printf("Fetching all products with the '%s' keyword from the '%s' category...", keyword, category)
	return s.Stock.getProductsByCategoryAndKeyword(keyword, category)
}
