package query

import (
	"loja/internal/inventory/adapter/output/model/db"
	"loja/internal/inventory/application/dto"
	"loja/internal/configuration/handler_err"

	"gorm.io/gorm"
)

type GetProducts struct {
	db *gorm.DB
}

func NewQueryGetProducts(db *gorm.DB) *GetProducts {
	return &GetProducts{
		db: db,
	}
}

func (gp *GetProducts) Run(inventoryInput dto.GetProductsInput) ([]dto.GetProductsOutput, *handler_err.InfoErr) {
	var products []db.ProductInventoryDB
	if err := gp.db.Where("seller_id = ?", inventoryInput.SellerID).Find(&products).Error; err != nil {
		return []dto.GetProductsOutput{}, &handler_err.InfoErr{
			Message: "seller not found",
			Err: handler_err.ErrNotFound,
		}
	}

	var productsOutput []dto.GetProductsOutput
	for _, product := range products {
		productInfo := dto.GetProductsOutput{
			ID: product.ID,
			Description: product.Description,
			Price: product.Price,
			Quantity: product.Quantity,
			SellerID: product.SellerID,
		}

		productsOutput = append(productsOutput, productInfo)
	}

	return productsOutput, &handler_err.InfoErr{}
}