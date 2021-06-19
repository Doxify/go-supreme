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

// Represents metadata of a product
type Data struct {
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
