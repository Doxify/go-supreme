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
// Create a new instance and fetch the latest stock from supreme.
s := CreateNewInstance()
s.FetchStock()

// Get all products in the "hats" category.
p := s.Stock.GetProducts("hats")
```

## III. Features
This project is currently at a very early stage and a work in progress. More
features will be added soon.

1. Fetch latest stock from Supreme
2. Get products by category
3. Get products by keyword


## IV. Todo

**Planned:**
- ⚪ Keep track of changes in stock
- ⚪ Fetch stock at an interval
- ⚪ Monitor stock by product keyword

**In-progress:**
- 🟡 Add product styles




