package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

// connect to the Db
func init() {
	var err error
	Db, err = sql.Open("mysql", "edukan:edukan@/edukaan")
	if err != nil {
		panic(err)
	}
}

// Get a single vendor.  mySQL uses ? and not $1 - try the below and refer to
func retrieve(id int) (vendor Vendor, err error) {
	vendor = Vendor{}
	err = Db.QueryRow("select id, `name`, owner, address from vendor where id = ?", id).Scan(&vendor.Id, &vendor.Name, &vendor.Owner, &vendor.Address)
	return
}

// Create a new vendor
func (vendor *Vendor) create() (id int64, err error) {
	statement := "insert into vendor (name, owner, address) values (?, ?, ?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(vendor.Name, vendor.Owner, vendor.Address)
	return result.LastInsertId()

}

// Update a vendor
func (vendor *Vendor) update() (err error) {
	_, err = Db.Exec("update vendor set `name` = ?, owner = ? , address = ? where id = ?", vendor.Name, vendor.Owner, vendor.Address, vendor.Id)
	if err != nil {
		panic(err)
	}
	return
}

// Delete a vendor
func (vendor *Vendor) delete() (err error) {
	_, err = Db.Exec("delete from vendor where id = ?", vendor.Id)
	return
}
