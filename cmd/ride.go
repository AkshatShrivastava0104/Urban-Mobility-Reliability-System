package main

import(
	"fmt"
	"github.com/gorilla/mux"
)

type Ride struct{
	id int
	rider_id int
	driver_id int
	vehicle_id int
	pickup_location string
	dropoff_location string
	status bool
	requested_at string
	completed_at string
}


func main(){
	// ride := mux.NewRouter();
	// ride.HandleFunc()
}