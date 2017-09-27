package models

type Customer struct {
	Id           int
	Name         string
	MobileNumber int
	Address      string
}

type Vendor struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Owner   string `json:"owner"`
	Address string `json:"address"`
}

type Order struct {
	Id           int
	CustomerId   int
	VendorId     int
	OrderDetails string
	Status       string
}
