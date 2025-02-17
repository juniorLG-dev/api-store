package query

import (
	"loja/internal/configuration/handler_err"
	"loja/internal/seller/application/dto"
	"loja/internal/seller/adapter/output/model/db"

	"gorm.io/gorm"
)

type GetSellerByUsername struct {
	db *gorm.DB
}

func NewQueryGetSellerByUsername(db *gorm.DB) *GetSellerByUsername {
	return &GetSellerByUsername{
		db: db,
	}
}

func (gs *GetSellerByUsername) Run(sellerInput dto.GetSellerByUsernameInput) (*dto.GetSellerByUsernameOutput, *handler_err.InfoErr) {
	var seller db.SellerDB
	if err := gs.db.First(&seller, "username = ?", sellerInput.Username).Error; err != nil {
		return &dto.GetSellerByUsernameOutput{}, &handler_err.InfoErr{
			Message: "seller not found",
			Err: handler_err.ErrNotFound,
		}
	}

	return &dto.GetSellerByUsernameOutput{
		ID: seller.ID,
		Name: seller.Name,
		Username: seller.Username,
		Email: seller.Email,
	}, &handler_err.InfoErr{}
}