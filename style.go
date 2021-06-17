package gosupreme

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

type Size struct {
	Name       string `json:"name"`
	ID         int    `json:"id"`
	StockLevel int    `json:"stock_level"`
}

type Sizes []*Size
