package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/doxify/supreme-go-api/supreme"
)

func main() {
	s := supreme.CreateNewInstance()
	s.FetchStock()

	p := s.LookupCategoryProduct("panel", "hats")
	for _, product := range *p {
		spew.Dump(product)
	}
}
