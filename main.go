package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ovidiugiorgi/wsproduct/handler"
	"github.com/ovidiugiorgi/wsproduct/store"
	"github.com/urfave/negroni"
)

func main() {
	n := negroni.Classic()

	r := mux.NewRouter()
	// var c = handler.NewProductController(store.NewInMemStore())
	var c = handler.NewProductController(store.NewRedisStore())
	registerProductRoutes(r, c)

	n.UseFunc(jsonMw)
	n.UseHandler(r)

	log.Fatal(http.ListenAndServe(":8080", n))
}
