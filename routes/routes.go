package routes

import (
	"github.com/gorilla/mux"
	"github.com/jerevick8/handlers"
	"log"
	"net/http"
	"os"
)

var L = log.New(os.Stdout, "Product API", log.LstdFlags)

func Routes() http.Handler {
	serveMux := mux.NewRouter()

	//hh := handlers.NewHello(l)
	ah := handlers.NewAbout(L)
	ph := handlers.NewProducts(L)

	//Routes for GET Requests

	getRoutes := serveMux.Methods(http.MethodGet).Subrouter()
	getRoutes.HandleFunc("/about", ah.About)
	getRoutes.HandleFunc("/products", ph.GetProducts)

	//Routes for POST Requests
	postRoutes := serveMux.Methods(http.MethodPost).Subrouter()
	postRoutes.HandleFunc("/products", ph.AddProducts)

	//Routes for PUT Requests
	putRoutes := serveMux.Methods(http.MethodPut).Subrouter()
	putRoutes.HandleFunc("/{id:[0-9]+}", ph.UpdateProduct)
	//serveMux.Handle("/", hh)
	return serveMux
}
