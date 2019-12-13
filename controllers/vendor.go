package controllers

import (
	"edukaan/common"
	"edukaan/errors"
	"edukaan/models"
	"edukaan/repository"
	"encoding/json"
	"net/http"
)

type VendorController struct {
	repo repository.VendorRepository
}

var Vendor VendorController

func init() {
	Vendor := VendorController{}
	Vendor.repo = repository.VendorRepository{}
}

func (controller *VendorController) FindVendor(w http.ResponseWriter, r *http.Request) (error *errors.AppError) {

	name := r.URL.Query().Get("name")
	common.Info.Println("Search query name ", name)

	vendors, err := controller.repo.FindVendors(name)
	if err != nil {
		common.Error.Println("Could not find any vendors with name: ", name, err)
		return &errors.AppError{Error: err, Message: "Order Id not provided", Code: -1}
	}
	output, err := json.MarshalIndent(&vendors, "", "\t\t")
	if err != nil {
		return &errors.AppError{Error: err, Message: "Order Id not provided", Code: -1}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(output)

	return nil
}

func (controller *VendorController) RetrieveVendor(writer http.ResponseWriter, r *http.Request) (error *errors.AppError) {

	id, error, hasError := ExtractPathParam(r, "id")
	if hasError {
		return error
	}
	vendor, err := controller.repo.Retrieve(id)
	if err != nil {
		common.Error.Println("Invalid vendor id", err)
		return &errors.AppError{Error: err, Message: "Invalid vendor Id", Code: 400}
	}
	common.Info.Println("vendor found %d", id)
	return WriteOKResponse(vendor, writer)
}

func (controller *VendorController) DeleteVendor(w http.ResponseWriter, r *http.Request) (error *errors.AppError) {

	id, error, hasError := ExtractPathParam(r, "id")
	if hasError {
		return error
	}

	vendor := models.Vendor{Id: id}
	err := controller.repo.Delete(&vendor)
	if err != nil {
		common.Error.Println("Vendor could not be deleted ", id)
		return
	}
	common.Info.Println("Vendor deleted ", id)
	w.WriteHeader(http.StatusOK)
	return

}

func (controller *VendorController) UpdateVendor(w http.ResponseWriter, r *http.Request) (error *errors.AppError) {
	var vendor models.Vendor
	id, error, done := ExtractIdentifierAndEntity(r, &vendor, w)
	if done {
		return error
	}

	vendor.Id = id
	err := controller.repo.Update(&vendor)

	if error != nil {
		common.Error.Println("Vendor could not be updated", err)
		return &errors.AppError{Error: err, Message: "Order Id not provided", Code: -1}
	}
	w.WriteHeader(http.StatusOK)
	// location header should also be set as per the REST standards
	return
}

func (controller *VendorController) CreateVendor(w http.ResponseWriter, r *http.Request) (error *errors.AppError) {

	var vendor models.Vendor
	appError, hasError := ParseRequest(r, &vendor, w)
	if hasError {
		return appError
	}

	id, err := controller.repo.Create(&vendor)
	if err != nil {
		common.Error.Println("Vendor could not be created", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	common.Info.Println("Created vendor with id ", id)
	return

}
