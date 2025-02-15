package usecase

import (
	"loja/internal/seller/adapter/output/repository"
	"loja/internal/configuration/handler_err"
	"loja/internal/seller/application/dto"
	"loja/internal/common/domain/entities"
	"loja/internal/common/domain/service"
)

type LoginSeller struct {
	repository repository.PortRepository
}

func NewUseCaseLoginSeller(repository repository.PortRepository) *LoginSeller {
	return &LoginSeller{
		repository: repository,
	}
}

func (ls *LoginSeller) Run(sellerInput dto.LoginSellerInput) (string, *handler_err.InfoErr) {
	sellerDomain, err := ls.repository.GetSellerByEmail(sellerInput.Email)
	if err != nil {
		return "", &handler_err.InfoErr{
			Message: "email or password incorrect",
			Err: handler_err.ErrInvalidInput,
		}
	}

	if !sellerDomain.Password.CheckPassword(sellerInput.Password) {
		return "", &handler_err.InfoErr{
			Message: "email or password incorrect",
			Err: handler_err.ErrInvalidInput,
		}
	}

	user, msgErr := entities.NewUser(
		sellerDomain.ID,
		sellerDomain.Name,
		sellerDomain.Username.Value,
	)

	if msgErr.Err != nil {
		return "", msgErr
	}

	tokenGenerator := service.NewTokenGenerator("seller") 

	token, msgErr := tokenGenerator.GenerateToken(user)
	if msgErr.Err != nil {
		return "", msgErr
	}

	return token, &handler_err.InfoErr{}
}