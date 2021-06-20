package main

import (
	"log"

	gosupreme "github.com/doxify/gosupreme"
)

func main() {
	s := gosupreme.New()
	s.Init()

	s.GetProductsByCategory(gosupreme.ProductCategory.Accessories)

	p, err := s.GetProductsByKeyword("Brooks")
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, product := range *p {
		s.FetchProductData(product)
		s.FetchProductData(product)
	}

}
