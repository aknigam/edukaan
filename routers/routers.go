package routers

import (
	"edukaan/common"
	"edukaan/controllers"
	"edukaan/errors"
	"github.com/gorilla/mux"
	"net/http"
)

func SetVendorRoutes(r *mux.Router) {

	r.Handle("/vendors/{id}", appHandler{controllers.Vendor.RetrieveVendor}).Methods("GET")
	// find the vendor
	r.Handle("/vendors", appHandler{controllers.Vendor.FindVendor}).Methods("GET")
	r.Handle("/vendors/{id}", appHandler{controllers.Vendor.UpdateVendor}).Methods("PUT")
	r.Handle("/vendors", appHandler{controllers.Vendor.CreateVendor}).Methods("POST")
	r.Handle("/vendors/{id}", appHandler{controllers.Vendor.DeleteVendor}).Methods("DELETE")

	r.Handle("/orders/{id}", appHandler{controllers.Order.RetrieveOrder}).Methods("GET")
	r.Handle("/orders/{id}", appHandler{controllers.Order.UpdateOrder}).Methods("PUT")
	r.Handle("/orders", appHandler{controllers.Order.CreateOrder}).Methods("POST")
	r.Handle("/orders/{id}", appHandler{controllers.Order.DeleteOrder}).Methods("DELETE")
	// list the orders for a vendor /orders/vendor/{id}?status=NEW|Delivered|Out for delivery
	r.Handle("/orders/vendor/{id}", appHandler{controllers.Order.FindVendorOrders}).Methods("GET")

	r.Handle("/customers/{id}", appHandler{controllers.Customer.RetrieveCustomer}).Methods("GET")
	r.Handle("/customers/{id}", appHandler{controllers.Customer.UpdateCustomer}).Methods("PUT")
	r.Handle("/customers", appHandler{controllers.Customer.CreateCustomer}).Methods("POST")
	r.Handle("/customers/{id}", appHandler{controllers.Customer.DeleteCustomer}).Methods("DELETE")

	// accept the order /orders/{id}/respond?accept=true|false
	r.HandleFunc("/orders/{id}/respond", DoNothing).Methods("POST")

	// check order status
	r.HandleFunc("/orders/{id}/status", DoNothing).Methods("GET")

	// change order status
	r.HandleFunc("/orders/{id}/statusupdate", DoNothing).Methods("POST")

}

func DoNothing(writer http.ResponseWriter, r *http.Request) {

}

type appHandler struct {
	Handler func(http.ResponseWriter, *http.Request) *errors.AppError
}

// https://blog.golang.org/error-handling-and-go
func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn.Handler(w, r); e != nil { // e is *appError, not os.Error.
		common.Error.Println("Unexpected error", e)
		//http.Error(writer, err.Error(), http.StatusBadRequest)
		http.Error(w, e.Message, e.Code)
		return
	}
}
