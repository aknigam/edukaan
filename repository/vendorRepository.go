package repository

import (
	"edukaan/common"
	"edukaan/models"
	_ "github.com/go-sql-driver/mysql"
	"log"
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
	err = Db.QueryRow("select id, `name`, owner, address from vendor where id = ?", id).
		Scan(&vendor.Id, &vendor.Name, &vendor.Owner, &vendor.Address)
	if err != nil {
		txn.Rollback()
	}
	return vendor, err
}

// Create a new vendor
func (repo *VendorRepository) Create(vendor *models.Vendor) (id int, err error) {
	statement := "insert into vendor (name, owner, address) values (?, ?, ?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(vendor.Name, vendor.Owner, vendor.Address)
	generatedId, err := result.LastInsertId()
	return int(generatedId), err

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

// TODO - incomplete implementation
func (repo *VendorRepository) FindVendors(name string) (s []models.Vendor, err error) {

	rows, err := Db.Query("select id, `name`, owner, address from vendor where `name` like %?% ", name)
	if err != nil {
		common.Error.Println("Could not find orders ", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		vendor := models.Vendor{}
		err := rows.Scan(&vendor.Id, &vendor.Name, &vendor.Owner, &vendor.Address)
		if err != nil {
			common.Error.Println("Could not find vendors ", err)
			break
		}
		s = append(s, vendor)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return
}
