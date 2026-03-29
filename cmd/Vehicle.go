package main

import(
	"fmt"
	"github.com/gorilla/mux"
)

type Vehicle struct{
	ID int
	driver_id int
	vehicle_type string
	capacity int
	wheelchair_support bool
	child_seats int
}


func main(){
	// ride := mux.NewRouter();
	// ride.HandleFunc()
}