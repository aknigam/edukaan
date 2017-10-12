package repository

import (
	"edukaan/common"
	"edukaan/models"
	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepository struct {
}

/*
 Get a single customer.  mySQL uses ? and not $1 - try the below and refer to
 Transaction management code is only for demonstration. it can saftely be removed.
*/
func (repo *CustomerRepository) Retrieve(id int) (customer models.Customer, err error) {

	customer = models.Customer{}
	if err != nil {
		return
	}
	err = Db.QueryRow("select id, `name`, address, mobile from customer where id = ?", id).
		Scan(&customer.Id, &customer.Name, &customer.Address, &customer.MobileNumber)
	return customer, err
}

// Create a new customer
func (repo *CustomerRepository) Create(customer *models.Customer) (id int64, err error) {
	statement := "insert into customer (name, address, mobile) values (?, ?, ?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		common.Error.Println("Customer could not be created ", err)
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(customer.Name, customer.Address, customer.MobileNumber)
	if err != nil {
		common.Error.Println("Customer could not be created ", err)
		return
	}
	id, err = result.LastInsertId()
	if err != nil {
		common.Error.Println("Customer could not be created ", err)
		return
	}
	return
}

// Update a customer
func (repo *CustomerRepository) Update(customer *models.Customer) (err error) {
	_, err = Db.Exec("update customer set `name` = ?, mobile = ? , address = ? where id = ?",
		customer.Name, customer.MobileNumber, customer.Address, customer.Id)
	if err != nil {
		common.Error.Println("Customer could not be updated ")
		return
	}
	return
}

// Delete a customer
func (repo *CustomerRepository) Delete(customer *models.Customer) (err error) {
	_, err = Db.Exec("delete from customer where id = ?", customer.Id)
	return
}
