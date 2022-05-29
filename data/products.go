package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	SKU         string  `json:"-"`
	Price       float64 `json:"price"`
	CreatedOn   string  `json:"_"`
	DeleteOn    string  `json:"_"`
	UpdatedOn   string  `json:"_"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) FromJSON(w io.Reader) error {
	e := json.NewDecoder(w)
	return e.Decode(p)
}
func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}
func GetProducts() Products {
	return productList
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "LG Television",
		Description: "An ultra-wide screen with 4K resolution display",
		Price:       100.00,
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "IPhone 14",
		Description: "High-resolution camera",
		Price:       50.00,
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
