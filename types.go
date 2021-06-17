package gosupreme

//  Wrapper that sorts all products by their categories.
type Categories struct {
	Accessories  Products `json:"Accessories"`
	Bags         Products `json:"Bags"`
	Skate        Products `json:"Skate"`
	TopsSweaters Products `json:"Tops/Sweaters"`
	Pants        Products `json:"Pants"`
	Jackets      Products `json:"Jackets"`
	Hats         Products `json:"Hats"`
	Sweatshirts  Products `json:"Sweatshirts"`
	Shirts       Products `json:"Shirts"`
	Shorts       Products `json:"Shorts"`
	New          Products `json:"new"`
}
