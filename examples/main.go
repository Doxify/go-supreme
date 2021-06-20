package main

import (
	"fmt"
	"log"

	gosupreme "github.com/doxify/gosupreme"
)

func main() {
	s := gosupreme.New()
	s.Init()

	p, err := s.GetProductsByKeyword("Boxer Briefs (2 Pack)")
	if err != nil {
		log.Fatal(err)
	}

	for _, product := range *p {
		s.FetchProductData(product)
		styles, _ := s.GetAllStyles(product)
		style, _ := styles.GetStyleByColor("purple")
		size, _ := style.GetSize("medium")

		fmt.Println(product)
		fmt.Println(style)
		fmt.Println(size)

	}

}
