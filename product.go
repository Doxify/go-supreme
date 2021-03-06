package gosupreme

import (
	"fmt"
	"reflect"
	"strings"
)

type Product struct {
	Name         string `json:"name"`
	ID           int    `json:"id"`
	ImageURL     string `json:"image_url"`
	ImageURLHi   string `json:"image_url_hi"`
	Price        int    `json:"price"`
	SalePrice    int    `json:"sale_price"`
	NewItem      bool   `json:"new_item"`
	Position     int    `json:"position"`
	CategoryName string `json:"category_name"`
}

type Products []*Product

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetID() int {
	return p.ID
}

func (p *Product) GetPrice() int {
	// Supreme formats their price weirdly
	return p.Price / 100
}

func (p *Product) GetCategory() string {
	return p.CategoryName
}

func (p *Product) String() string {
	return fmt.Sprintf("Product: Name: %s\nID: %d\nCategory: %s\n Price:%d\n", p.GetName(), p.GetID(), p.GetCategory(), p.GetPrice())
}

// Returns a list of all products in a category.
func (s *Supreme) GetProductsByCategory(category string) (*Products, error) {

	fields := reflect.TypeOf(s.Stock.Categories)
	values := reflect.ValueOf(s.Stock.Categories)
	num := fields.NumField()

	// loop through all categories
	for i := 0; i < num; i++ {
		c := fields.Field(i)
		p := values.Field(i)

		// if category exists, return its products
		if strings.EqualFold(category, c.Name) {
			products, _ := p.Interface().(Products)
			return &products, nil
		}
	}

	// return nil if category does not exist
	return nil, fmt.Errorf("invalid category; %s is not a valid category", category)
}

func (s *Supreme) GetProductsByKeyword(keyword string) (*Products, error) {

	if s.Stock == nil {
		return nil, fmt.Errorf("stock is nil; fetch stock before calling this function")
	}

	var matchedProducts Products

	fields := reflect.TypeOf(s.Stock.Categories)
	values := reflect.ValueOf(s.Stock.Categories)
	num := fields.NumField()

	// loop through all categories
	for i := 0; i < num; i++ {
		p := values.Field(i)
		products, _ := p.Interface().(Products)

		// loop through all products in a category
		for _, product := range products {
			// find products which have the keyword in their names
			if strings.Contains(strings.ToLower(product.GetName()), strings.ToLower(keyword)) {
				matchedProducts = append(matchedProducts, product)
			}
		}

	}

	if len(matchedProducts) == 0 {
		return nil, fmt.Errorf("no products found with keyword: %s", keyword)
	}

	return &matchedProducts, nil
}

func (s *Supreme) GetProductsByKeywordAndCategory(keyword string, category string) (*Products, error) {
	// storing all products that match the keyword and category
	var matchedProducts Products

	// getting all products that match the category
	products, err := s.GetProductsByCategory(category)
	if err != nil {
		return nil, err
	}

	// finding products in the category that match the keyword
	for _, p := range *products {
		if strings.Contains(strings.ToLower(p.GetName()), strings.ToLower(keyword)) {
			matchedProducts = append(matchedProducts, p)
		}
	}

	return &matchedProducts, nil
}
