package usecase

import (
	"loja/internal/seller/adapter/output/repository"
	"loja/internal/configuration/handler_err"
	"loja/internal/seller/application/dto"
)

type GetSellerByUsername struct {
	repository repository.PortRepository
}

func NewUseCaseGetSellerByUsername(repository repository.PortRepository) *GetSellerByUsername {
	return &GetSellerByUsername{
		repository: repository,
	}
}

func (gs *GetSellerByUsername) Run(sellerInput dto.GetSellerByUsernameInput) (dto.GetSellerByUsernameOutput, *handler_err.InfoErr) {
	seller, err := gs.repository.GetSellerByUsername(sellerInput.Username)
	if err != nil {
		return dto.GetSellerByUsernameOutput{}, &handler_err.InfoErr{
			Message: "seller not found",
			Err: handler_err.ErrNotFound,
		}
	}

	sellerOutput := dto.GetSellerByUsernameOutput{
		ID: seller.ID.Value,
		Name: seller.Name,
		Username: seller.Username.Value,
		Email: seller.Email,
	}

	return sellerOutput, &handler_err.InfoErr{}
}