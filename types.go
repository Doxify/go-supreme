package gosupreme

//  Wrapper that sorts all products by their categories.
type Categories struct {
	Accessories  Products `json:"Accessories"`
	Bags         Products `json:"Bags"`
	Skate        Products `json:"Skate"`
	TopsSweaters Products `json:"Tops/Sweaters"`
	Pants        Products `json:"Pants"`
	Jackets      Products `json:"Jackets"`
	Hats         Products `json:"Hats"`
	Sweatshirts  Products `json:"Sweatshirts"`
	Shirts       Products `json:"Shirts"`
	Shorts       Products `json:"Shorts"`
	New          Products `json:"new"`
}

// Represents a product's style
type Style struct {
	ID                    int         `json:"id"`
	Name                  string      `json:"name"`
	Chk                   string      `json:"chk"`
	Tag                   interface{} `json:"tag"`
	Currency              string      `json:"currency"`
	Description           string      `json:"description"`
	ImageURL              string      `json:"image_url"`
	ImageURLHi            string      `json:"image_url_hi"`
	SwatchURL             string      `json:"swatch_url"`
	SwatchURLHi           string      `json:"swatch_url_hi"`
	MobileZoomedURL       string      `json:"mobile_zoomed_url"`
	MobileZoomedURLHi     string      `json:"mobile_zoomed_url_hi"`
	BiggerZoomedURL       string      `json:"bigger_zoomed_url"`
	SpecialPurchasableQty interface{} `json:"special_purchasable_qty"`
	Sizes                 Sizes       `json:"sizes"`
}

type Styles []*Style

// Represents a product style size
type Size struct {
	Name       string `json:"name"`
	ID         int    `json:"id"`
	StockLevel int    `json:"stock_level"`
}

type Sizes []*Size
