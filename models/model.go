package models

type Customer struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	MobileNumber int    `json:"mobile"`
	Address      string `json:"address"`
}

type Vendor struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Owner   string `json:"owner"`
	Address string `json:"address"`
}

type Order struct {
	Id           int    `json:"id"`
	CustomerId   int    `json:"customerId"`
	VendorId     int    `json:"vendorId"`
	OrderDetails string `json:"orderDetails"`
	Status       string `json:"status"`
}
