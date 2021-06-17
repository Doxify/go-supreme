package gosupreme

import (
	"log"
	"net/http"
	"os"
	"time"
)

type Supreme struct {
	// Stock holds Supreme's product inventory
	Stock *Stock

	// Styles holds Supreme's product styles
	Styles *Styles

	l *log.Logger
	c *http.Client
}

// Creates a new gosupreme instance
func New() *Supreme {
	return &Supreme{
		Stock:  nil,
		Styles: nil,
		l:      log.New(os.Stdout, "gosupreme ", log.LstdFlags),
		c: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (s *Supreme) Init() {
	s.FetchStock()
}

// package gosupreme

// func (s *Supreme) GetAllCategoryProducts(category string) *Products {
// 	s.Client.Logger.Printf("Fetching all products from the '%s' category...", category)
// 	return s.Stock.getProductsByCategory(category)
// }

// func (s *Supreme) GetAllProducts(keyword string) *Products {
// 	s.Client.Logger.Printf("Fetching all products with the '%s' keyword...", keyword)
// 	return s.Stock.getProductsByKeyword(keyword)
// }

// func (s *Supreme) LookupCategoryProduct(keyword string, category string) *Products {
// 	s.Client.Logger.Printf("Fetching all products with the '%s' keyword from the '%s' category...", keyword, category)
// 	return s.Stock.getProductsByCategoryAndKeyword(keyword, category)
// }
