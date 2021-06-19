package gosupreme

import (
	"fmt"
	"net/http"
)

// Represents the a product style
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

func (d *Data) String() string {
	return fmt.Sprintf("Description: %s\n", d.Description)
}

func (s *Supreme) FetchProductData(p *Product) (*Data, error) {
	s.l.Printf("Fetching product data for %d (%s)\n", p.ID, p.Name)

	// create a new http GET request
	req, err := http.NewRequest("GET", fmt.Sprintf(DataURL, p.ID), nil)
	if err != nil {
		return nil, err
	}

	// make the http request and save its response
	var data *Data
	err = s.makeRequest(req, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
