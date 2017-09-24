package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/vendors/{id}", retrieveVendor).Methods("GET")
	r.HandleFunc("/vendors/{id}", updateVendor).Methods("PUT")
	r.HandleFunc("/vendors", createVendor).Methods("POST")
	r.HandleFunc("/vendors/{id}", deleteVendor).Methods("DELETE")

	http.Handle("/", r)
	server := http.Server{
		Addr: ":8080",
	}
	server.ListenAndServe()

}
