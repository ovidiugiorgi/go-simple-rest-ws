package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	registerProductRoutes(r)
	r.Use(commonMw)
	r.Use(loggingMw)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
