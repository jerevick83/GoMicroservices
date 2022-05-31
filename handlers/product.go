package handlers

import (
	"github.com/jerevick8/data"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l: l}
}
func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		p.getProducts(w, r)
	}
	if r.Method == "POST" {
		p.addProducts(w, r)
	}
}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	lp := data.GetProducts()
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal product", http.StatusInternalServerError)
	}
}

func (p *Products) addProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Post Product")
	newProduct := &data.Product{}
	err := newProduct.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "unable to unmarshal product", http.StatusBadRequest)
	}
	data.AddProduct(newProduct)
	p.getProducts(w, r)
}
