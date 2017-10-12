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

type OrderController struct {
	repo repository.OrderRepository
}

var Order OrderController

func init() {
	Order := OrderController{}
	Order.repo = repository.OrderRepository{}
}

func (controller *OrderController) RetrieveOrder(writer http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(path.Base(vars["id"]))
	if err != nil {
		common.Error.Println("Invalid order id", err)
		return
	}
	order, err := controller.repo.Retrieve(id)
	if err != nil {
		common.Error.Println("Invalid order id", err)
		return
	}
	common.Info.Println("order found %d", id)
	output, err := json.MarshalIndent(&order, "", "\t\t")
	if err != nil {
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(output)

	return
}

func (controller *OrderController) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(path.Base(vars["id"]))
	if err != nil {
		return
	}

	order := models.Order{Id: id}
	controller.repo.Delete(&order)
	if err != nil {
		common.Error.Println("Order could not be deleted ", id)
		return
	}
	common.Info.Println("Order deleted ", id)
	w.WriteHeader(http.StatusOK)

}

func (controller *OrderController) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(path.Base(vars["id"]))
	if err != nil {
		return
	}

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var order models.Order
	json.Unmarshal(body, &order)
	order.Id = id
	error := controller.repo.Update(&order)
	if error != nil {
		common.Error.Println("Order could not be updated", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	// location header should also be set as per the REST standards
	return
}

func (controller *OrderController) CreateOrder(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var order models.Order
	json.Unmarshal(body, &order)

	id, err := controller.repo.Create(&order)
	if err != nil {
		common.Error.Println("Order could not be created", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	common.Info.Println("Created order with id ", id)
	return
}

func (controller *OrderController) FindVendorOrders(w http.ResponseWriter, r *http.Request) *error {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(path.Base(vars["id"]))
	if err != nil {
		common.Error.Println("Invalid vendor id", err)
		return &err

	}
	orders, err := controller.repo.FindOrders(int64(id))
	if err != nil {
		common.Error.Println("Invalid order id", err)
		return &err
	}
	common.Info.Println("order found %d", id)
	output, err := json.MarshalIndent(&orders, "", "\t\t")
	if err != nil {
		return &err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(output)

	return nil
}
