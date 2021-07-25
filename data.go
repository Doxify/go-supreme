package gosupreme

import (
	"fmt"
	"net/http"
)

// Wrapper for a Product's metadata
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

func (d *Data) String() string {
	return fmt.Sprintf("Description: %s\n", d.Description)
}

func (s *Supreme) FetchProductData(p *Product) (*Data, error) {
	// Return cached version if it exists
	if s.Data[p.ID] != nil {
		return s.Data[p.ID], nil
	}

	s.l.Printf("Fetching product data for %d (%s)\n", p.ID, p.Name)

	// create a new http GET request
	req, err := http.NewRequest("GET", fmt.Sprintf(DATA_URL, p.ID), nil)
	if err != nil {
		return nil, err
	}

	// make the http request and save its response
	var data *Data
	err = s.makeRequest(req, &data)
	if err != nil {
		return nil, err
	}

	// Store in cache
	s.Data[p.ID] = data

	return data, nil
}
