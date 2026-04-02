package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func health(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/health", health)

	SetupDriverRoutes(router)
	SetupRiderRoutes(router)
	SetupVehicleRoutes(router)
	SetupRideRoutes(router)
	SetupLocationRoutes(router)

	fmt.Println("Starting Server on port 8080...")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
