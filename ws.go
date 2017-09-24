package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"path"
	"strconv"

)

type Vendor struct {
	Id      int    `json:"id"`
	Name string `json:"name"`
	Owner  string `json:"owner"`
}

func retrieveVendor(writer http.ResponseWriter, request *http.Request) (err error ){

	id, err := strconv.Atoi(path.Base(request.URL.Path))
	if err != nil {
		return
	}
	vendor := Vendor{
		Id: id,
		Name: "medical shop",
		Owner: "Anand",

	}
	output, err := json.MarshalIndent(&vendor, "", "\t\t")
	if err != nil {
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(output)

	return
}

func deleteVendor(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

func updateVendor(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

func createVendor(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}
func main() {

	http.HandleFunc("/vendor/", handleRequest)
	server := http.Server{
		Addr: ":8080",
	}
	server.ListenAndServe()

}
func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = retrieveVendor(w, r)

	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
