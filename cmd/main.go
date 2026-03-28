package main

import (
	"fmt"
	"net/http"
)

func health(w http.ResponseWriter, req *http.Request){
	w.WriteHeader(http.StatusOK);
	w.Write([]byte("OK"));
}


func Ride()

func driver()

func events()

func ETA()

func SLA()

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintln(w, "Hello, your Go server is working!")
	// })

	http.HandleFunc("/health", health);

	fmt.Println("Starting Server on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}