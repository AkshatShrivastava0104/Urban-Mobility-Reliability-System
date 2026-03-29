package main

import(
	"fmt"
	"github.com/gorilla/mux"
)

type Driver struct{
	ID string 
	Name string
	rating float32
	status bool
	driver_location string
}


func main(){
	// ride := mux.NewRouter();
	// ride.HandleFunc()
}