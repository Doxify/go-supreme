package main

import gosupreme "github.com/doxify/supreme-go-api"

func main() {
	s := gosupreme.New()
	s.FetchStock()
}
