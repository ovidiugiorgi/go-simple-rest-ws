package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ovidiugiorgi/wsproduct/handler"
)

func createHandler(fn func(http.ResponseWriter, *http.Request) error) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := fn(w, r)
		if err != nil {
			log.Printf("error: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte(err.Error()))
			if err != nil {
				panic(err)
			}
		}
	}
}

func registerProductRoutes(r *mux.Router, c *handler.ProductController) {
	s := r.PathPrefix("/products").Subrouter()
	s.HandleFunc("", createHandler(c.CreateProduct)).Methods("POST")
	s.HandleFunc("", createHandler(c.ListProducts)).Methods("GET")
	s.HandleFunc("/{productID}", createHandler(c.GetProduct)).Methods("GET")
	s.HandleFunc("/{productID}", createHandler(c.RemoveProduct)).Methods("DELETE")
}
