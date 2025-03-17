package query

import (
	"loja/internal/customer/application/dto"
	"loja/internal/customer/adapter/output/model/db"
	"loja/internal/configuration/handler_err"

	"gorm.io/gorm"
)

type GetCustomerByID struct {
	db *gorm.DB
}

func NewQueryGetCustomerByID(db *gorm.DB) *GetCustomerByID {
	return &GetCustomerByID{
		db: db,
	}
}

func (gc *GetCustomerByID) Run(customerInput dto.GetCustomerByIDInput) (dto.GetCustomerByIDOutput, *handler_err.InfoErr) {
	var customer db.CustomerDB
	if err := gc.db.First(&customer, "id = ?", customerInput.ID).Error; err != nil {
		return dto.GetCustomerByIDOutput{}, &handler_err.InfoErr{
			Message: "customer not found",
			Err: handler_err.ErrNotFound,
		}
	}

	return dto.GetCustomerByIDOutput{
		ID: customer.ID,
		Name: customer.Name,
		Username: customer.Username,
		Email: customer.Email,
	}, &handler_err.InfoErr{}
}