package repository

import (
	"edukaan/common"
	"edukaan/models"
	_ "github.com/go-sql-driver/mysql"
)

type OrderRepository struct {
}

/*
 Get a single order.  mySQL uses ? and not $1 - try the below and refer to
 Transaction management code is only for demonstration. it can saftely be removed.
*/
func (repo *OrderRepository) Retrieve(id int) (order models.Order, err error) {
	order = models.Order{}
	txn, err := Db.Begin()
	if err != nil {
		return
	}
	defer txn.Commit()
	err = Db.QueryRow("select id, `name`, owner, address from order where id = ?", id).Scan(&order.Id, &order.CustomerId, &order.CustomerId, &order.CustomerId)
	if err != nil {
		txn.Rollback()
	}
	return order, err
}

// Create a new order
func (repo *OrderRepository) Create(order *models.Order) (id int64, err error) {
	statement := "insert into order (name, owner, address) values (?, ?, ?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(order.CustomerId, order.CustomerId, order.CustomerId)
	return result.LastInsertId()

}

// Update a order
func (repo *OrderRepository) Update(order *models.Order) (err error) {
	_, err = Db.Exec("update order set `name` = ?, owner = ? , address = ? where id = ?", order.CustomerId, order.CustomerId, order.CustomerId, order.Id)
	if err != nil {
		common.Error.Println("Order could not be updated ")
		panic(err)
	}
	return
}

// Delete a order
func (repo *OrderRepository) Delete(order *models.Order) (err error) {
	_, err = Db.Exec("delete from order where id = ?", order.Id)
	return
}
