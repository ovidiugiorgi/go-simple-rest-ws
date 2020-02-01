package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ovidiugiorgi/wsproduct/handler"
	"github.com/ovidiugiorgi/wsproduct/model"
	"github.com/ovidiugiorgi/wsproduct/store"
	"github.com/urfave/negroni"
)

func getStorage(flag string) model.ProductService {
	switch flag {
	case "redis":
		return store.NewRedisStore()
	case "in-memory":
		return store.NewInMemStore()
	default:
		return store.NewRedisStore()
	}
}

func main() {
	storageFlag := flag.String(
		"storage",
		"redis",
		"storage backend. Options: \"redis\", \"in-memory\"",
	)
	portFlag := flag.Int(
		"port",
		8080,
		"API port",
	)
	flag.Parse()

	n := negroni.Classic()
	r := mux.NewRouter()

	c := handler.NewProductController(getStorage(*storageFlag))
	registerProductRoutes(r, c)

	n.UseFunc(jsonMw)
	n.UseHandler(r)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*portFlag), n))
}
