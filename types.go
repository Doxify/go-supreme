package gosupreme

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
