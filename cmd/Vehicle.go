package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Vehicle struct {
	ID                 int
	driver_id          int
	vehicle_type       string
	capacity           int
	wheelchair_support bool
	child_seats        int
}

func SetupVehicleRoutes(router *mux.Router) {
	router.HandleFunc("/vehicle", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Vehicle Route is Ready"))
	}).Methods("GET")
}
