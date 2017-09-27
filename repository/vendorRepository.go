package repository

import (
	"edukaan/common"
	"edukaan/models"
	_ "github.com/go-sql-driver/mysql"
)

type VendorRepository struct {
}

/*
 Get a single vendor.  mySQL uses ? and not $1 - try the below and refer to
 Transaction management code is only for demonstration. it can saftely be removed.
*/
func (repo *VendorRepository) Retrieve(id int) (vendor models.Vendor, err error) {
	vendor = models.Vendor{}
	txn, err := Db.Begin()
	if err != nil {
		return
	}
	defer txn.Commit()
	err = Db.QueryRow("select id, `name`, owner, address from vendor where id = ?", id).Scan(&vendor.Id, &vendor.Name, &vendor.Owner, &vendor.Address)
	if err != nil {
		txn.Rollback()
	}
	return vendor, err
}

// Create a new vendor
func (repo *VendorRepository) Create(vendor *models.Vendor) (id int64, err error) {
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
func (repo *VendorRepository) Update(vendor *models.Vendor) (err error) {
	_, err = Db.Exec("update vendor set `name` = ?, owner = ? , address = ? where id = ?", vendor.Name, vendor.Owner, vendor.Address, vendor.Id)
	if err != nil {
		common.Error.Println("Vendor could not be updated ")
		panic(err)
	}
	return
}

// Delete a vendor
func (repo *VendorRepository) Delete(vendor *models.Vendor) (err error) {
	_, err = Db.Exec("delete from vendor where id = ?", vendor.Id)
	return
}
