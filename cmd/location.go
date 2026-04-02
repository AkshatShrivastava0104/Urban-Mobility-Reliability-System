package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type location struct {
	longitude string
	latitude  string
}

func SetupLocationRoutes(router *mux.Router) {
	router.HandleFunc("/location", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Location Route is Ready"))
	}).Methods("GET")
}
