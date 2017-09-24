package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"path"
	"strconv"
)

type Vendor struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Owner   string `json:"owner"`
	Address string `json:"address"`
}

func retrieveVendor(writer http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(path.Base(vars["id"]))
	if err != nil {
		Error.Println("Invalid vendor id")
		return
	}
	vendor, err := retrieve(id)
	if err != nil {
		Error.Println("Invalid vendor id")
		return
	}
	Info.Println("vendor found %d", id)
	output, err := json.MarshalIndent(&vendor, "", "\t\t")
	if err != nil {
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(output)

	return
}

func deleteVendor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(path.Base(vars["id"]))
	if err != nil {
		return
	}
	vendor := Vendor{Id: id}
	vendor.delete()
	if err != nil {
		Error.Println("Vendor could not be deleted ", id)
		return
	}
	Info.Println("Vendor deleted ", id)
	w.WriteHeader(http.StatusOK)

}

func updateVendor(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var vendor Vendor
	json.Unmarshal(body, &vendor)
	err := vendor.update()
	if err != nil {
		Error.Println("Vendor could not be update", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}

func createVendor(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var vendor Vendor
	json.Unmarshal(body, &vendor)
	id, err := vendor.create()
	if err != nil {
		Error.Println("Vendor could not be created", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	Info.Println("Created vendor with id ", id)
	return

}
