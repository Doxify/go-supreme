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

**Output of the code above:**
```shell
gosupreme 2021/06/18 19:54:53.722725 Fetching latest stock from supreme.
gosupreme 2021/06/18 19:54:54.072299 Fetched stock with release date '06/17/2021'.
Name: Warp Jacquard Logos Denim 6-Panel
ID: 174325
Category: Hats
gosupreme 2021/06/18 19:54:54.072342 Fetching product data for 174325 (Warp Jacquard Logos Denim 6-Panel)
Description: All cotton 14 oz. denim 6-Panel hat with jacquard logo pattern and self strap closure. Embroidered logo on front.

Name: Liberty Floral 6-Panel
ID: 174278
Category: Hats
gosupreme 2021/06/18 19:54:54.120455 Fetching product data for 174278 (Liberty Floral 6-Panel)
Description: All cotton 6-Panel hat with self strap closure and embroidered logo on front.

Name: Vampire Boy 6-Panel
ID: 174284
Category: Hats
gosupreme 2021/06/18 19:54:54.208470 Fetching product data for 174284 (Vampire Boy 6-Panel)
Description: Wool blend 6-Panel hat with snap closure. Embroidered graphic on front and embroidered logo on back. Original artwork by Sean Cliver.

Name: Warp Jacquard Logos Denim 6-Panel
ID: 174325
Category: Hats
gosupreme 2021/06/18 19:54:54.251057 Fetching product data for 174325 (Warp Jacquard Logos Denim 6-Panel)
Description: All cotton 14 oz. denim 6-Panel hat with jacquard logo pattern and self strap closure. Embroidered logo on front.

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
