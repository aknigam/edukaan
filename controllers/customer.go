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

type CustomerController struct {
	repo repository.CustomerRepository
}

var Customer CustomerController

func init() {
	Customer := CustomerController{}
	Customer.repo = repository.CustomerRepository{}
}

func (controller *CustomerController) RetrieveCustomer(writer http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(path.Base(vars["id"]))
	if err != nil {
		common.Error.Println("Invalid customer id", err)
		return
	}
	customer, err := controller.repo.Retrieve(id)
	if err != nil {
		common.Error.Println("Invalid customer id", err)
		return
	}
	common.Info.Println("customer found %d", id)
	output, err := json.MarshalIndent(&customer, "", "\t\t")
	if err != nil {
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(output)

	return
}

func (controller *CustomerController) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(path.Base(vars["id"]))
	if err != nil {
		return
	}

	customer := models.Customer{Id: id}
	controller.repo.Delete(&customer)
	if err != nil {
		common.Error.Println("Customer could not be deleted ", id)
		return
	}
	common.Info.Println("Customer deleted ", id)
	w.WriteHeader(http.StatusOK)

}

func (controller *CustomerController) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(path.Base(vars["id"]))
	if err != nil {
		return
	}

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var customer models.Customer
	json.Unmarshal(body, &customer)
	customer.Id = id
	error := controller.repo.Update(&customer)
	if error != nil {
		common.Error.Println("Customer could not be updated", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	// location header should also be set as per the REST standards
	return
}

func (controller *CustomerController) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var customer models.Customer
	json.Unmarshal(body, &customer)

	id, err := controller.repo.Create(&customer)
	if err != nil {
		common.Error.Println("Customer could not be created", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	common.Info.Println("Created customer with id ", id)
	return

}
