package query

import (
	"loja/internal/cart/adapter/output/model/db"
	"loja/internal/cart/application/dto"
	"loja/internal/common/domain/service"
	"loja/internal/configuration/handler_err"


	"gorm.io/gorm"
)

type GetProducts struct {
	db *gorm.DB
	tokenGenerator *service.TokenGenerator
}

func NewQueryGetProducts(db *gorm.DB) *GetProducts {
	tokenGenerator := service.NewTokenGenerator("customer")
	return &GetProducts{
		db: db,
		tokenGenerator: tokenGenerator,
	}
}

func (gp *GetProducts) Run(token string) ([]dto.GetProductsOutput, *handler_err.InfoErr) {
	customer, msgErr := gp.tokenGenerator.VerifyToken(token)
	if msgErr.Err != nil {
		return []dto.GetProductsOutput{}, msgErr
	}

	if customer.Type != gp.tokenGenerator.Value {
		return []dto.GetProductsOutput{}, &handler_err.InfoErr{
			Message: "you don't have a cart",
			Err: handler_err.ErrUnauthorized,
		}
	}
	var products []db.CartDB
	err := gp.db.Where("customer_id = ?", customer.ID).Find(&products).Error
	if err != nil {
		return []dto.GetProductsOutput{}, &handler_err.InfoErr{
			Message: "empty cart",
			Err: handler_err.ErrNotFound,
		}
	}

	var productsOutput []dto.GetProductsOutput
	for _, product := range products {
		productInfo := dto.GetProductsOutput{
			ID: product.ID,
			Description: product.Description,
			Price: product.Price,
			ProductID: product.ProductID,
			CustomerID: product.CustomerID,
			SellerID: product.SellerID,
		}

		productsOutput = append(productsOutput, productInfo)
	}

	return productsOutput, &handler_err.InfoErr{}
}