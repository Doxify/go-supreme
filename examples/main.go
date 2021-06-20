package main

import (
	"log"

	gosupreme "github.com/doxify/supreme-go-api"
)

func main() {
	s := gosupreme.New()
	s.Init()

	s.GetProductsByCategory(gosupreme.ProductCategory.Accessories)

	p, err := s.GetProductsByKeyword("6-panel")
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, product := range *p {
		// fmt.Print(product)
		s.FetchProductData(product)
		s.FetchProductData(product)
	}

}
