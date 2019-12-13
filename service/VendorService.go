package service

import (
	"edukaan/errors"
	"edukaan/models"
	"edukaan/repository"
)

type VendorService struct {
	Repo      repository.VendorRepository
	OrderRepo repository.OrderRepository
}

func (service *VendorService) Retrieve(id int) (vendor models.Vendor, appErr *errors.AppError) {
	vendor, err := service.Repo.Retrieve(id)
	if err != nil {
		appErr = &errors.AppError{Error: err, Message: "vendor not found", Code: -1}
	}
	return vendor, appErr
}

func (service *VendorService) Delete(vendor *models.Vendor) (appErr *errors.AppError) {
	err := service.Repo.Delete(vendor)
	if err != nil {
		appErr = &errors.AppError{Error: err, Message: "vendor could not be deleted", Code: -1}
	}
	return appErr
}

func (service *VendorService) Update(vendor *models.Vendor) (appErr *errors.AppError) {
	err := service.Repo.Update(vendor)
	if err != nil {
		appErr = &errors.AppError{Error: err, Message: "vendor could not be updated", Code: -1}
	}
	return appErr
}

func (service *VendorService) Create(vendor *models.Vendor) (id int, appErr *errors.AppError) {
	id, err := service.Repo.Create(vendor)
	if err != nil {
		appErr = &errors.AppError{Error: err, Message: "vendor could not be created", Code: -1}
	}
	if vendor.Orders != nil {
		for _, v := range vendor.Orders {
			_, err = service.OrderRepo.Create(&v)
			if err != nil {
				appErr = &errors.AppError{Error: err, Message: "order could not be created", Code: -1}
				break
			}
		}
	}
	return id, appErr
}

func (service *VendorService) FindVendors(vendorId int64) (vendors []models.Vendor, appErr *errors.AppError) {
	vendors, err := service.Repo.FindVendors("vendorId")
	if err != nil {
		appErr = &errors.AppError{Error: err, Message: "vendors not found", Code: -1}
	}
	return vendors, appErr
}
