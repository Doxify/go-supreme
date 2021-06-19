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
// Create a new instance and initialize it.
s := gosupreme.New()
s.Init()

// Get all products that have the keyword "6-panel" in their name
p, err := s.GetProductsByKeyword("6-panel")
if err != nil {
  log.Fatal(err)
  return
}

// Loop through all of those products and fetch their product data (styles, sizes, etc...) and print it to console.
for _, product := range *p {
  fmt.Print(product)
  d, err := s.FetchProductData(product)
  if err != nil {
    log.Fatal(err)
    return
  }
  fmt.Println(d)
}

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
