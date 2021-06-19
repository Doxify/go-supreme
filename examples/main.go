package main

import (
	"fmt"
	"log"

	gosupreme "github.com/doxify/supreme-go-api"
)

func main() {
	s := gosupreme.New()
	s.Init()

	p, err := s.GetProductsByKeyword("6-panel")
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, product := range *p {
		fmt.Print(product)
		d, err := s.FetchProductData(product)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println(d)
	}

}
