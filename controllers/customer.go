package controllers

import (
	"edukaan/common"
	"edukaan/errors"
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

// https://blog.golang.org/error-handling-and-go
func (controller *CustomerController) RetrieveCustomer(writer http.ResponseWriter, r *http.Request) (error *errors.AppError) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(path.Base(vars["id"]))
	if err != nil {
		common.Error.Println("Invalid customer id", err)
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	customer, err := controller.repo.Retrieve(id)
	if err != nil {

		common.Error.Println("Invalid customer id", err)
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	common.Info.Println("customer found %d", id)
	output, err := json.MarshalIndent(&customer, "", "\t\t")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(output)

	return
}

func (controller *CustomerController) DeleteCustomer(w http.ResponseWriter, r *http.Request) (error *errors.AppError) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(path.Base(vars["id"]))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	customer := models.Customer{Id: id}
	controller.repo.Delete(&customer)
	if err != nil {
		common.Error.Println("Customer could not be deleted ", id)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return &errors.AppError{err, "Order Id not provided", -1}
	}
	common.Info.Println("Customer deleted ", id)
	w.WriteHeader(http.StatusOK)
	return
}

func (controller *CustomerController) UpdateCustomer(w http.ResponseWriter, r *http.Request) (error *errors.AppError) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(path.Base(vars["id"]))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var customer models.Customer
	json.Unmarshal(body, &customer)
	customer.Id = id
	err = controller.repo.Update(&customer)
	if error != nil {
		common.Error.Println("Customer could not be updated", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return &errors.AppError{err, "Order Id not provided", -1}
	}
	w.WriteHeader(http.StatusOK)
	// location header should also be set as per the REST standards
	return
}

func (controller *CustomerController) CreateCustomer(w http.ResponseWriter, r *http.Request) (error *errors.AppError) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var customer models.Customer
	json.Unmarshal(body, &customer)

	id, err := controller.repo.Create(&customer)
	if err != nil {
		common.Error.Println("Customer could not be created", err)
		return &errors.AppError{err, "Order Id not provided", -1}
	}
	w.WriteHeader(http.StatusOK)
	common.Info.Println("Created customer with id ", id)
	return nil

}
