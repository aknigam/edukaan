package controllers

import (
	"edukaan/common"
	"edukaan/errors"
	"edukaan/models"
	"edukaan/repository"
	"edukaan/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"path"
	"strconv"
)

type OrderController struct {
	orderService service.OrderService
	//repo repository.OrderRepository
}

var Order OrderController

func init() {
	Order := OrderController{}
	Order.orderService = service.OrderService{
		Repo: repository.OrderRepository{},
	}
}

// refer: https://blog.golang.org/error-handling-and-go
func (controller *OrderController) RetrieveOrder(writer http.ResponseWriter, r *http.Request) (error *errors.AppError) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(path.Base(vars["id"]))
	if err != nil {
		common.Error.Println("Invalid order id", err)
		return &errors.AppError{err, "Order Id not provided", -1}
	}
	order, error := controller.orderService.Retrieve(id)
	if err != nil {
		common.Error.Println("Invalid order id", err)
		return error
	}
	common.Info.Println("order found %d", id)
	// use encoder here instead of marshallking . That will be 50% faster
	// https://learning.oreilly.com/library/view/building-microservices-with/9781786468666/4f3ac8c1-a27f-4d1d-819d-2a16e51bb7b3.xhtml
	// /Users/a.nigam/Documents/workspace-old/gospace/src/BuildingMicroserviceswithGo_Code/Chapter01/chapter1-master/reading_writing_json_3/reading_writing_json_3.go
	output, err := json.MarshalIndent(&order, "", "\t\t")
	if err != nil {
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(output)

	return
}

func (controller *OrderController) DeleteOrder(w http.ResponseWriter, r *http.Request) (error *errors.AppError) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(path.Base(vars["id"]))
	if err != nil {
		return &errors.AppError{err, "Order Id not provided", -1}
	}

	order := models.Order{Id: id}
	error = controller.orderService.Delete(&order)
	if error != nil {
		return error
	}
	common.Info.Println("Order deleted ", id)
	w.WriteHeader(http.StatusOK)

	return
}

func (controller *OrderController) UpdateOrder(w http.ResponseWriter, r *http.Request) (error *errors.AppError) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(path.Base(vars["id"]))
	if err != nil {
		return
	}

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var order models.Order
	// instead of unmarshal use decoder
	// refer: Chapter01/chapter1-master/reading_writing_json_5/reading_writing_json_5.go
	json.Unmarshal(body, &order)
	order.Id = id
	error = controller.orderService.Update(&order)
	if error != nil {
		common.Error.Println("Order could not be updated", err)
		return error
	}
	w.WriteHeader(http.StatusOK)
	// location header should also be set as per the REST standards
	return
}

func (controller *OrderController) CreateOrder(w http.ResponseWriter, r *http.Request) (error *errors.AppError) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var order models.Order
	json.Unmarshal(body, &order)

	id, err := controller.orderService.Create(&order)
	if err != nil {
		common.Error.Println("Order could not be created", err)
		return err
	}
	w.WriteHeader(http.StatusOK)
	common.Info.Println("Created order with id ", id)
	return
}

func (controller *OrderController) FindVendorOrders(w http.ResponseWriter, r *http.Request) (error *errors.AppError) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(path.Base(vars["id"]))
	if err != nil {
		common.Error.Println("Invalid vendor id", err)
		return &errors.AppError{err, "Order Id not provided", -1}
	}
	orders, error := controller.orderService.FindOrders(int64(id))
	if err != nil {
		return error
	}
	common.Info.Println("orders not found ")
	output, marshalError := json.MarshalIndent(&orders, "", "\t\t")
	if err != nil {
		return &errors.AppError{marshalError, "Order Id not provided", -1}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(output)

	return nil
}
