package controllers

import (
	"edukaan/common"
	"edukaan/models"
	"edukaan/repository"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"path"
	"strconv"
)

type VendorController struct {
	repo repository.VendorRepository
}

var Vendor VendorController

func init() {
	Vendor := VendorController{}
	Vendor.repo = repository.VendorRepository{}
}

func (controller *VendorController) FindVendor(w http.ResponseWriter, r *http.Request) *error {

	name := r.URL.Query().Get("name")
	common.Info.Println("Search query name ", name)

	vendors, err := controller.repo.FindVendors(name)
	if err != nil {
		common.Error.Println("Could not find any vendors with name: ", name, err)
		return &err
	}
	output, err := json.MarshalIndent(&vendors, "", "\t\t")
	if err != nil {
		return &err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(output)

	return nil
	return nil
}

func (controller *VendorController) RetrieveVendor(writer http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(path.Base(vars["id"]))
	if err != nil {
		common.Error.Println("Invalid vendor id", err)
		return
	}
	vendor, err := controller.repo.Retrieve(id)
	if err != nil {
		common.Error.Println("Invalid vendor id", err)
		return
	}
	common.Info.Println("vendor found %d", id)
	output, err := json.MarshalIndent(&vendor, "", "\t\t")
	if err != nil {
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(output)

	return
}

func (controller *VendorController) DeleteVendor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(path.Base(vars["id"]))
	if err != nil {
		return
	}

	vendor := models.Vendor{Id: id}
	controller.repo.Delete(&vendor)
	if err != nil {
		common.Error.Println("Vendor could not be deleted ", id)
		return
	}
	common.Info.Println("Vendor deleted ", id)
	w.WriteHeader(http.StatusOK)

}

func (controller *VendorController) UpdateVendor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(path.Base(vars["id"]))
	if err != nil {
		return
	}

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var vendor models.Vendor
	json.Unmarshal(body, &vendor)
	vendor.Id = id
	error := controller.repo.Update(&vendor)
	if error != nil {
		common.Error.Println("Vendor could not be updated", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	// location header should also be set as per the REST standards
	return
}

func (controller *VendorController) CreateVendor(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var vendor models.Vendor
	json.Unmarshal(body, &vendor)

	id, err := controller.repo.Create(&vendor)
	if err != nil {
		common.Error.Println("Vendor could not be created", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	common.Info.Println("Created vendor with id ", id)
	return

}
