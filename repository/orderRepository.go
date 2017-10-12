package repository

import (
	"edukaan/common"
	"edukaan/models"
	_ "github.com/go-sql-driver/mysql"

	"log"
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
	err = Db.QueryRow("select id, `vendorId`, customerId, orderDetails, status from `order` where id = ?", id).
		Scan(&order.Id, &order.VendorId, &order.CustomerId, &order.OrderDetails, &order.Status)
	if err != nil {
		txn.Rollback()
	}
	return order, err
}

// Create a new order
func (repo *OrderRepository) Create(order *models.Order) (id int64, err error) {
	statement := "insert into `order` (vendorId, customerId, orderDetails, status) values (?, ?, ?, ?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(order.VendorId, order.CustomerId, order.OrderDetails, order.Status)
	if err != nil {
		common.Error.Println("Order could not be created ", err)
		return
	}
	return result.LastInsertId()

}

// Update a order
func (repo *OrderRepository) Update(order *models.Order) (err error) {
	_, err = Db.Exec("update `order` set orderDetails = ? where id = ?", order.OrderDetails, order.Id)
	if err != nil {
		common.Error.Println("Order could not be updated ")
		panic(err)
	}
	return
}

// Delete a order
func (repo *OrderRepository) Delete(order *models.Order) (err error) {
	_, err = Db.Exec("delete from `order` where id = ?", order.Id)
	return
}

// refer: http://go-database-sql.org/retrieving.html
func (repo *OrderRepository) FindOrders(vendorId int64) (s []models.Order, err error) {
	rows, err := Db.Query("select id, `vendorId`, customerId, orderDetails, status from `order` where vendorId = ?", vendorId)
	if err != nil {
		common.Error.Println("Could not find orders ", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		order := models.Order{}
		err := rows.Scan(&order.Id, &order.VendorId, &order.CustomerId, &order.OrderDetails, &order.Status)
		if err != nil {
			common.Error.Println("Could not find orders ", err)
			break
		}
		s = append(s, order)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return
}
