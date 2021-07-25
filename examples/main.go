package main

import (
	"flag"
	"fmt"

	gosupreme "github.com/doxify/gosupreme"
)

func main() {
	// Parse command line flags
	keywordPtr := flag.String("keyword", "", "Keyword that must be in the product name")
	colorPtr := flag.String("color", "", "Color of a product")
	sizePtr := flag.String("size", "", "Size of a product")
	flag.Parse()

	// Create a new instance of gosupreme and initialize it.
	s := gosupreme.New()
	s.Init()

	// Get all products that have the keyword 'Boxer Briefs (2 Pack)' in their
	// name.
	products, err := s.GetProductsByKeyword(*keywordPtr)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Select the first one from the list of returned products
	var product = (*products)[0]

	// Fetch that products data (styles, sizes, etc...)
	s.FetchProductData(product)

	// Get all of the product's styles (these are the different colors/variants
	// the products come in)
	styles, _ := s.GetAllStyles(product)

	// Get a style that is the color 'purple'
	style, _ := styles.GetStyleByColor(*colorPtr)

	// Get the size that is 'medium'
	size, _ := style.GetSize(*sizePtr)

	// Print the product, style, and size to console.
	fmt.Println(product)
	fmt.Println(style)
	fmt.Println(size)
}
