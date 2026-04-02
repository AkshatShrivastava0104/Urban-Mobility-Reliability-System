package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Ride struct {
	id               int
	rider_id         int
	driver_id        int
	vehicle_id       int
	pickup_location  string
	dropoff_location string
	status           bool
	requested_at     string
	completed_at     string
}

func SetupRideRoutes(router *mux.Router) {
	router.HandleFunc("/ride", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Ride Route is Ready"))
	}).Methods("GET")
}

// func main(){
// 	// ride := mux.NewRouter();
// 	// ride.HandleFunc()
// }
