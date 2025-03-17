package usecase

import (
	"loja/internal/seller/adapter/output/repository"
	"loja/internal/seller/application/dto"
	"loja/internal/configuration/handler_err"
	"loja/internal/common/domain/service"
)

type DeleteSeller struct {
	repository repository.PortRepository
	tokenGenerator *service.TokenGenerator
}

func NewUseCaseDeleteSeller(repo repository.PortRepository) *DeleteSeller {
	tokenGenerator := service.NewTokenGenerator("seller")
	return &DeleteSeller{
		repository: repo,
		tokenGenerator: tokenGenerator,
	}
}

func (ds *DeleteSeller) Run(sellerInput dto.DeleteSellerInput, token string) *handler_err.InfoErr {
	seller, msgErr := ds.tokenGenerator.VerifyToken(token)
	if msgErr.Err != nil {
		return msgErr
	}

	sellerDomain, err := ds.repository.GetSellerByID(seller.ID)
	if err != nil {
		return &handler_err.InfoErr{
			Message: "could not find this seller",
			Err: handler_err.ErrInternal,
		}
	}

	if !sellerDomain.Password.CheckPassword(sellerInput.Password) {
		return &handler_err.InfoErr{
			Message: "invalid password",
			Err: handler_err.ErrInvalidInput,
		}
	}

	if err = ds.repository.DeleteSeller(seller.ID); err != nil {
		return &handler_err.InfoErr{
			Message: "could not delete user",
			Err: handler_err.ErrInternal,
		}
	}

	return &handler_err.InfoErr{}
}