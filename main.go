package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	n := negroni.Classic()

	r := mux.NewRouter()
	registerProductRoutes(r)

	n.UseFunc(jsonMw)
	n.UseHandler(r)

	log.Fatal(http.ListenAndServe(":8080", n))
}
