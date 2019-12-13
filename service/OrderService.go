package service

import (
	"edukaan/errors"
	"edukaan/models"
	"edukaan/repository"
)

type OrderService struct {
	Repo repository.OrderRepository
}

func (service *OrderService) Retrieve(id int) (order models.Order, appErr *errors.AppError) {
	order, err := service.Repo.Retrieve(id)
	if err != nil {
		appErr = &errors.AppError{Error: err, Message: "order not found", Code: -1}
	}
	return order, appErr
}

func (service *OrderService) Delete(order *models.Order) (appErr *errors.AppError) {
	err := service.Repo.Delete(order)
	if err != nil {
		appErr = &errors.AppError{Error: err, Message: "order could not be deleted", Code: -1}
	}
	return appErr
}

func (service *OrderService) Update(order *models.Order) (appErr *errors.AppError) {
	err := service.Repo.Update(order)
	if err != nil {
		appErr = &errors.AppError{Error: err, Message: "order could not be updated", Code: -1}
	}
	return appErr
}

func (service *OrderService) Create(order *models.Order) (id int, appErr *errors.AppError) {
	id, err := service.Repo.Create(order)
	if err != nil {
		appErr = &errors.AppError{Error: err, Message: "order could not be created", Code: -1}
	}
	return id, appErr
}

func (service *OrderService) FindOrders(vendorId int64) (orders []models.Order, appErr *errors.AppError) {
	orders, err := service.Repo.FindOrders(vendorId)
	if err != nil {
		appErr = &errors.AppError{Error: err, Message: "orders not found", Code: -1}
	}
	return orders, appErr
}
