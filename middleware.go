package main

import (
	"net/http"
)

func jsonMw(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Header().Add("Content-Type", "application/json")
	next.ServeHTTP(w, r)
}
