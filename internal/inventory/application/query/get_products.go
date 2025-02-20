package query

import (
	"loja/internal/inventory/application/dto"
	"loja/internal/configuration/handler_err"
	"loja/internal/inventory/adapter/output/model/db"

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

func (gp *GetProducts) Run(sellerInput dto.GetProductsInput) ([]dto.GetProductsOutput, *handler_err.InfoErr) {
	var products []db.ProductInventoryDB
	err := gp.db.Where("seller_id = ?", sellerInput.SellerID).Find(&products).Error
	if err != nil {
		return []dto.GetProductsOutput{}, &handler_err.InfoErr{
			Message: "seller not found",
			Err: handler_err.ErrNotFound,
		}
	}

	var productsInventory []dto.GetProductsOutput
	for _, product := range products {
		productInventoryInfo := dto.GetProductsOutput{
			ID: product.ID,
			Description: product.Description,
			Price: product.Price,
			Quantity: product.Quantity,
		}

		productsInventory = append(productsInventory, productInventoryInfo)
	}

	return productsInventory, &handler_err.InfoErr{}
}