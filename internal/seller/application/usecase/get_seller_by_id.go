package usecase

import (
	"loja/internal/configuration/handler_err"
	"loja/internal/seller/adapter/output/repository"
	"loja/internal/seller/application/dto"
)

type GetSellerByID struct {
	repository repository.PortRepository
}

func NewUseCaseGetSellerByID(repo repository.PortRepository) *GetSellerByID {
	return &GetSellerByID{
		repository: repo,
	}
}

func (gs *GetSellerByID) Run(sellerInput dto.GetSellerByIDInput) (*dto.GetSellerByIDOutput, *handler_err.InfoErr) {
	seller, err := gs.repository.GetSellerByID(sellerInput.ID)
	if err != nil {
		return nil, &handler_err.InfoErr{
			Message: "user not found",
			Err: handler_err.ErrNotFound,
		}
	}

	sellerOutput := &dto.GetSellerByIDOutput{
		ID: seller.ID.Value,
		Name: seller.Name,
		Username: seller.Username.Value,
		Email: seller.Email,
	}

	return sellerOutput, nil
}