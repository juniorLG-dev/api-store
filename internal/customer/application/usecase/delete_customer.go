package usecase

import (
	"loja/internal/customer/adapter/output/repository"
	"loja/internal/customer/application/dto"
	"loja/internal/common/domain/service"
	"loja/internal/configuration/handler_err"
)

type DeleteCustomer struct {
	repository repository.PortRepository
	tokenGenerator *service.TokenGenerator
}

func NewUseCaseDeleteCustomer(repo repository.PortRepository) *DeleteCustomer {
	tokenGenerator := service.NewTokenGenerator("customer")
	return &DeleteCustomer{
		repository: repo,
		tokenGenerator: tokenGenerator,
	}
}

func (dc *DeleteCustomer) Run(customerInput dto.DeleteCustomerInput, token string) *handler_err.InfoErr {
	customer, msgErr := dc.tokenGenerator.VerifyToken(token)
	if msgErr.Err != nil {
		return msgErr
	}

	customerDomain, err := dc.repository.GetCustomerByID(customer.ID)
	if err != nil {
		return &handler_err.InfoErr{
			Message: "could not find this seller",
			Err: handler_err.ErrInternal,
		}
	}

	if !customerDomain.Password.CheckPassword(customerInput.Password) {
		return &handler_err.InfoErr{
			Message: "invalid password",
			Err: handler_err.ErrInvalidInput,
		}
	}

	if err = dc.repository.DeleteCustomer(customer.ID); err != nil {
		return &handler_err.InfoErr {
			Message: "could not delete user",
			Err: handler_err.ErrInternal,
		}
	}

	return &handler_err.InfoErr{}
}