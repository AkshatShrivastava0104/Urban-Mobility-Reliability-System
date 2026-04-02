package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

type Driver struct {
	ID              string
	Name            string
	rating          float32
	status          bool
	driver_location string
}

func SetupDriverRoutes(router *mux.Router) {
	router.HandleFunc("/driver", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Driver is Online"))
	}).Methods("GET")
}