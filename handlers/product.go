package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jerevick8/data"
	"log"
	"net/http"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l: l}
}

func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	lp := data.GetProducts()
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal product", http.StatusInternalServerError)
	}
}

func (p *Products) AddProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Post Product")
	newProduct := &data.Product{}
	err := newProduct.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "unable to unmarshal product", http.StatusBadRequest)
	}
	err = newProduct.Validate()
	if err != nil {
		http.Error(w, fmt.Sprintf("Validation Error: %s", err), http.StatusBadRequest)
		return
	}
	data.AddProduct(newProduct)
	p.GetProducts(w, r)
}

func (p *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Request")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id from string to int", http.StatusBadRequest)
	}
	p.l.Println("ID: ", id)
	prod := &data.Product{}

	err = prod.FromJSON(r.Body)
	err = prod.Validate()
	if err != nil {
		http.Error(w, fmt.Sprintf("Validation Error: %s", err), http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, "unable to marshal product", http.StatusBadRequest)
	}
	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}
	p.GetProducts(w, r)
}
