package models

type Customer struct {
	//Unfortunately we can't, as in Go, lowercase properties are not exported, Marshal will ignore these and will not include them in the output.
	//
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

type helloWorldResponse struct {
	// change the output field to be "message"
	Message string `json:"message"`
	// do not output this field
	Author string `json:"-"`
	// do not output the field if the value is empty
	Date string `json:",omitempty"`
	// convert output to a string and rename "id"
	Id int `json:"id, string"`
}
