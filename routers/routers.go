package routers

import (
	"edukaan/common"
	"edukaan/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func SetVendorRoutes(r *mux.Router) {

	r.HandleFunc("/vendors/{id}", controllers.Vendor.RetrieveVendor).Methods("GET")
	// find the vendor
	r.Handle("/vendors", appHandler{controllers.Vendor.FindVendor}).Methods("GET")
	r.HandleFunc("/vendors/{id}", controllers.Vendor.UpdateVendor).Methods("PUT")
	r.HandleFunc("/vendors", controllers.Vendor.CreateVendor).Methods("POST")
	r.HandleFunc("/vendors/{id}", controllers.Vendor.DeleteVendor).Methods("DELETE")

	r.HandleFunc("/orders/{id}", controllers.Order.RetrieveOrder).Methods("GET")
	r.HandleFunc("/orders/{id}", controllers.Order.UpdateOrder).Methods("PUT")
	r.HandleFunc("/orders", controllers.Order.CreateOrder).Methods("POST")
	r.HandleFunc("/orders/{id}", controllers.Order.DeleteOrder).Methods("DELETE")
	// list the orders for a vendor /orders/vendor/{id}?status=NEW|Delivered|Out for delivery
	r.Handle("/orders/vendor/{id}", appHandler{controllers.Order.FindVendorOrders}).Methods("GET")

	r.HandleFunc("/customers/{id}", controllers.Customer.RetrieveCustomer).Methods("GET")
	r.HandleFunc("/customers/{id}", controllers.Customer.UpdateCustomer).Methods("PUT")
	r.Handle("/customers", appHandler{controllers.Customer.CreateCustomer}).Methods("POST")
	r.HandleFunc("/customers/{id}", controllers.Customer.DeleteCustomer).Methods("DELETE")

	// accept the order /orders/{id}/respond?accept=true|false
	r.HandleFunc("/orders/{id}/respond", DoNothing).Methods("POST")

	// check order status
	r.HandleFunc("/orders/{id}/status", DoNothing).Methods("GET")

	// change order status
	r.HandleFunc("/orders/{id}/statusupdate", DoNothing).Methods("POST")

	r.HandleFunc("/notes/save", controllers.Note.AddNote).Methods("POST")

}

func DoNothing(writer http.ResponseWriter, r *http.Request) {

}

type appHandler struct {
	Handler func(http.ResponseWriter, *http.Request) *error
}

// https://blog.golang.org/error-handling-and-go
func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	/*
		should create the business objects instead of passing w and r
	*/

	if e := fn.Handler(w, r); e != nil { // e is *appError, not os.Error.
		common.Error.Println("Invalid customer id", e)
		err := *e
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
