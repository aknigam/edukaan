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
	txn, err := Db.Begin()
	if err != nil {
		return
	}
	defer txn.Commit()
	err = Db.QueryRow("select id, `name`, owner, address from customer where id = ?", id).Scan(&customer.Id, &customer.Name, &customer.Name, &customer.Address)
	if err != nil {
		txn.Rollback()
	}
	return customer, err
}

// Create a new customer
func (repo *CustomerRepository) Create(customer *models.Customer) (id int64, err error) {
	statement := "insert into customer (name, owner, address) values (?, ?, ?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(customer.Name, customer.Name, customer.Address)
	return result.LastInsertId()

}

// Update a customer
func (repo *CustomerRepository) Update(customer *models.Customer) (err error) {
	_, err = Db.Exec("update customer set `name` = ?, owner = ? , address = ? where id = ?", customer.Name, customer.Name, customer.Address, customer.Id)
	if err != nil {
		common.Error.Println("Customer could not be updated ")
		panic(err)
	}
	return
}

// Delete a customer
func (repo *CustomerRepository) Delete(customer *models.Customer) (err error) {
	_, err = Db.Exec("delete from customer where id = ?", customer.Id)
	return
}
