package gosupreme

import (
	"fmt"
	"strings"
)

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

func (s *Style) String() string {
	return fmt.Sprintf("Style:\nID: %d\nColor: %s\n", s.ID, s.Name)
}

// Returns a product's styles
func (s *Supreme) GetAllStyles(p *Product) (*Styles, error) {
	// check if the style is cached
	if s.Data[p.ID] != nil {
		return &s.Data[p.ID].Styles, nil
	}

	//  otherwise fetch it
	pd, err := s.FetchProductData(p)
	if err != nil {
		return nil, err
	}

	return &pd.Styles, nil
}

func (s *Styles) GetStyleByColor(color string) (*Style, error) {
	for _, style := range *s {
		if strings.EqualFold(strings.ToLower(style.Name), strings.ToLower(color)) {
			return style, nil
		}
	}

	return nil, fmt.Errorf("product does not contain style for color '%s'", color)
}

// Gets a size of a product style
func (s *Style) GetSize(size string) (*Size, error) {
	// If a style only has one size, return it
	if len(s.Sizes) == 1 {
		return s.Sizes[0], nil
	}

	for _, si := range s.Sizes {
		if strings.EqualFold(strings.ToLower(si.Name), strings.ToLower(size)) {
			return si, nil
		}
	}

	return nil, fmt.Errorf("product does not contain size '%s' for the style '%s'", size, s.Name)
}

func (s *Size) String() string {
	return fmt.Sprintf("Size:\nName: %s\nID: %d\n", s.Name, s.ID)
}

// Returns all of the sizes available for a given style
func (s *Style) GetAllSizes() (*Sizes, error) {
	return &s.Sizes, nil
}
