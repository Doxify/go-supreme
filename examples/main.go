package main

import (
	"fmt"

	gosupreme "github.com/doxify/gosupreme"
)

func main() {
	s := gosupreme.New()
	s.Init()

	// Get all products that have the keyword 'Boxer Briefs (2 Pack)' in their
	// name.
	products, _ := s.GetProductsByKeyword("Boxer Briefs (2 Pack)")

	// Select the first one from the list of returned products
	var product = (*products)[0]

	// Fetch that products data (styles, sizes, etc...)
	s.FetchProductData(product)

	// Get all of the product's styles (these are the different colors/variants
	// the products come in)
	styles, _ := s.GetAllStyles(product)

	// Get a style that is the color 'purple'
	style, _ := styles.GetStyleByColor("purple")

	// Get the size that is 'medium'
	size, _ := style.GetSize("medium")

	fmt.Println(product)
	fmt.Println(style)
	fmt.Println(size)
}
