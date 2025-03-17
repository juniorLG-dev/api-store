package usecase

import (
	"loja/internal/common/domain/service"
	"loja/internal/common/domain/entities"
	"loja/internal/configuration/handler_err"
	"loja/internal/customer/adapter/output/repository"
	"loja/internal/customer/application/dto"
)

type LoginCustomer struct {
	repository repository.PortRepository
}

func NewUseCaseLoginCustomer(repo repository.PortRepository) *LoginCustomer {
	return &LoginCustomer{
		repository: repo,
	}
}

func (lc *LoginCustomer) Run(customerInput dto.LoginCustomerInput) (string, *handler_err.InfoErr) {
	customerDomain, err := lc.repository.GetCustomerByEmail(customerInput.Email)
	if err != nil {
		return "", &handler_err.InfoErr{
			Message: "email or password incorrect",
			Err: handler_err.ErrInvalidInput,
		}
	}

	if !customerDomain.Password.CheckPassword(customerInput.Password) {
		return "", &handler_err.InfoErr{
			Message: "email or password incorrect",
			Err: handler_err.ErrInvalidInput,
		}
	}

	user, msgErr := entities.NewUser(
		customerDomain.ID,
		customerDomain.Name,
		customerDomain.Username.Value,
	)
	if msgErr.Err != nil {
		return "", msgErr
	}

	tokenGenerator := service.NewTokenGenerator("customer")

	token, msgErr := tokenGenerator.GenerateToken(user)
	if msgErr.Err != nil {
		return "", msgErr
	}

	return token, &handler_err.InfoErr{}
}