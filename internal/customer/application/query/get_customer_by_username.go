package query

import (
	"loja/internal/customer/adapter/output/model/db"
	"loja/internal/customer/application/dto"
	"loja/internal/configuration/handler_err"

	"gorm.io/gorm"
)

type GetCustomerByUsername struct {
	db *gorm.DB
}

func NewQueryGetCustomerByUsername(db *gorm.DB) *GetCustomerByUsername {
	return &GetCustomerByUsername{
		db: db,
	}
}

func (gc *GetCustomerByUsername) Run(customerInput dto.GetCustomerByUsernameInput) (dto.GetCustomerByUsernameOutput, *handler_err.InfoErr) {
	var customer db.CustomerDB
	if err := gc.db.First(&customer, "username = ?", customerInput.Username).Error; err != nil {
		return dto.GetCustomerByUsernameOutput{}, &handler_err.InfoErr{
			Message: "customer not found",
			Err: handler_err.ErrNotFound,
		}
	}

	return dto.GetCustomerByUsernameOutput{
		ID: customer.ID,
		Name: customer.Name,
		Username: customer.Username,
		Email: customer.Email,
	}, &handler_err.InfoErr{}

}