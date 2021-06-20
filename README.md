<p align="center">
  <h2 align="center">supreme-go-api</h2>
  <p align="center"> 
    An unofficial go api for supremenewyork.com
  </p>
  <p align="center"> 
    <h3 align="center">⚠️ WORK IN PROGRESS ⚠️</h3>
  </p>
</p>

---
## I. Motivation
Just wanted to get more familiar with writing go code and since I had some prior
experience with Supreme's backend I could focus mainly on writing go.

## II. Usage
This project is currently not ready for release but if you'd like to mess around
with it here's how...


**Example usage:**
```go
// Create a new instance of gosupreme and initialize it.
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

	// Print the product, style, and size to console.
	fmt.Println(product)
	fmt.Println(style)
	fmt.Println(size)
```

**Output of the code above:**
```shell
[gosupreme] 2021/06/20 12:31:04.812178 Fetching latest stock from supreme.
[gosupreme] 2021/06/20 12:31:05.178247 Fetched stock with release date '06/17/2021'.
[gosupreme] 2021/06/20 12:31:05.178287 Fetching product data for 173930 (Supreme®/Hanes® Boxer Briefs (2 Pack))
Product: Name: Supreme®/Hanes® Boxer Briefs (2 Pack)
ID: 173930
Category: Accessories
 Price:28

Style:
ID: 29598
Color: Purple

Size:
Name: Medium
ID: 85957
```
## III. Features
This project is currently at a very early stage and a work in progress. More
features will be added soon.

1. Fetch latest stock from Supreme
2. Get products by category
3. Get products by keyword
4. Get all product styles


## IV. Todo

**Planned:**
- ⚪ Keep track of changes in stock
- ⚪ Fetch stock at an interval

**In-progress:**
- 🟡 Make the API easier to use
- 🟡 Monitor stock by product keyword
