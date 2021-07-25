package gosupreme

// Endpoints
const (
	STOCK_URL = "https://supremenewyork.com/mobile_stock.json"
	DATA_URL  = "https://supremenewyork.com/shop/%d.json"
)

// Product categories registry.
// This registry represents all valid supreme product categories.
var ProductCategory = newCategoryRegistry()

// Creates a new product registry with all valid product categories.
func newCategoryRegistry() *categoryRegistry {
	return &categoryRegistry{
		Accessories:  "accessories",
		Bags:         "bags",
		Skate:        "skate",
		TopsSweaters: "tops/sweaters",
		Pants:        "pants",
		Shirts:       "shirts",
		Shorts:       "shorts",
		Hats:         "hats",
		Sweatshirts:  "sweatshirts",
		Jackets:      "jackets",
		Shoes:        "shoes",
		new:          "new",
	}
}

type categoryRegistry struct {
	Accessories  string
	Bags         string
	Skate        string
	TopsSweaters string
	Pants        string
	Shirts       string
	Shorts       string
	Hats         string
	Sweatshirts  string
	Jackets      string
	Shoes        string
	new          string
}
