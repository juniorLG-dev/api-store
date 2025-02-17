package query

import (
	"loja/internal/configuration/handler_err"
	"loja/internal/seller/application/dto"
	"loja/internal/seller/adapter/output/model/db"

	"gorm.io/gorm"
)

type GetSellerByID struct {
	db *gorm.DB
}

func NewQueryGetSellerByID(db *gorm.DB) *GetSellerByID {
	return &GetSellerByID{
		db: db,
	}
}

func (gs *GetSellerByID) Run(sellerInput dto.GetSellerByIDInput) (*dto.GetSellerByIDOutput, *handler_err.InfoErr) {
	var seller db.SellerDB
	if err := gs.db.First(&seller, "id = ?", sellerInput.ID).Error; err != nil {
		return &dto.GetSellerByIDOutput{}, &handler_err.InfoErr{
			Message: "seller not found",
			Err: handler_err.ErrNotFound,
		}
	}

	return &dto.GetSellerByIDOutput{
		ID: seller.ID,
		Name: seller.Name,
		Username: seller.Username,
		Email: seller.Email,
	}, &handler_err.InfoErr{}
} 