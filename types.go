package gosupreme

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

type Style struct {
	Styles                  Styles      `json:"styles"`
	Description             string      `json:"description"`
	CanAddStyles            bool        `json:"can_add_styles"`
	CanBuyMultiple          bool        `json:"can_buy_multiple"`
	Ino                     string      `json:"ino"`
	CodBlocked              bool        `json:"cod_blocked"`
	CanadaBlocked           bool        `json:"canada_blocked"`
	MexicoBlocked           bool        `json:"mexico_blocked"`
	PurchasableQty          int         `json:"purchasable_qty"`
	NewItem                 bool        `json:"new_item"`
	Apparel                 bool        `json:"apparel"`
	Handling                int         `json:"handling"`
	NoFreeShipping          bool        `json:"no_free_shipping"`
	CanBuyMultipleWithLimit int         `json:"can_buy_multiple_with_limit"`
	Tag                     interface{} `json:"tag"`
}

type Styles []*Style
