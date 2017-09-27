package routers

import (
	"edukaan/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func SetVendorRoutes(r *mux.Router) {

	r.HandleFunc("/vendors/{id}", controllers.Vendor.RetrieveVendor).Methods("GET")
	r.HandleFunc("/vendors/{id}", controllers.Vendor.UpdateVendor).Methods("PUT")
	r.HandleFunc("/vendors", controllers.Vendor.CreateVendor).Methods("POST")
	r.HandleFunc("/vendors/{id}", controllers.Vendor.DeleteVendor).Methods("DELETE")

	// find the vendor
	r.HandleFunc("/vendors/find", DoNothing).Methods("POST")

	// place the order
	r.HandleFunc("/orders", DoNothing).Methods("POST")

	// list the orders for a vendor /orders/vendor/{id}?status=NEW|Delivered|Out for delivery
	r.HandleFunc("/orders/vendor/{id}", DoNothing).Methods("GET")

	// accept the order /orders/{id}/respond?accept=true|false
	r.HandleFunc("/orders/{id}/respond", DoNothing).Methods("POST")

	// check order status
	r.HandleFunc("/orders/{id}/status", DoNothing).Methods("GET")

	// change order status
	r.HandleFunc("/orders/{id}/statusupdate", DoNothing).Methods("POST")

	r.HandleFunc("/customers/{id}", controllers.Customer.RetrieveCustomer).Methods("GET")
	r.HandleFunc("/customers/{id}", controllers.Customer.UpdateCustomer).Methods("PUT")
	r.HandleFunc("/customers", controllers.Customer.CreateCustomer).Methods("POST")
	r.HandleFunc("/customers/{id}", controllers.Customer.DeleteCustomer).Methods("DELETE")

}

func DoNothing(writer http.ResponseWriter, r *http.Request) {

}
