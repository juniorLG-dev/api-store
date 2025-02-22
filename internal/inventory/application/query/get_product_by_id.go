package query

import (
	"loja/internal/inventory/adapter/output/model/db"
	"loja/internal/inventory/application/dto"
	"loja/internal/configuration/handler_err"

	"gorm.io/gorm"
)

type GetProductByID struct {
	db *gorm.DB
}

func NewQueryGetProductByID(db *gorm.DB) *GetProductByID {
	return &GetProductByID{
		db: db,
	}
}

func (gp *GetProductByID) Run(productInput dto.GetProductByIDInput) (dto.GetProductByIDOutput, *handler_err.InfoErr) {
	var product db.ProductInventoryDB
	err := gp.db.First(&product, "id = ?", productInput.ID).Error
	if err != nil {
		return dto.GetProductByIDOutput{}, &handler_err.InfoErr{
			Message: "product not found",
			Err: handler_err.ErrNotFound,
		}
	}

	return dto.GetProductByIDOutput{
		ID: product.ID,
		Description: product.Description,
		Price: product.Price,
		Quantity: product.Quantity,
		SellerID: product.SellerID,
	}, &handler_err.InfoErr{}
}