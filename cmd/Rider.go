package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Rider struct {
	ID           string
	Name         string
	rating       float32
	Phone_number int64
}

func SetupRiderRoutes(router *mux.Router) {

	router.HandleFunc("/rider", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Rider is Online"))
	}).Methods("GET")
}
