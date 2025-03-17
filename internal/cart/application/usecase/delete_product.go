package usecase

import (
	"loja/internal/cart/adapter/output/repository"
	"loja/internal/cart/application/dto"
	"loja/internal/common/domain/service"
	"loja/internal/configuration/handler_err"
)

type DeleteProduct struct {
	repository repository.PortRepository
	tokenGenerator *service.TokenGenerator
}

func NewUseCaseDeleteProduct(repo repository.PortRepository) *DeleteProduct {
	tokenGenerator := service.NewTokenGenerator("customer")
	return &DeleteProduct{
		repository: repo,
		tokenGenerator: tokenGenerator,
	}
}

func (dp *DeleteProduct) Run(cartInput dto.DeleteProductInput, token string) *handler_err.InfoErr {
	customer, msgErr := dp.tokenGenerator.VerifyToken(token)
	if msgErr.Err != nil {
		return msgErr
	}

	product, err := dp.repository.GetProductByID(cartInput.CartID)
	if err != nil {
		return &handler_err.InfoErr{
			Message: "product not found",
			Err: handler_err.ErrNotFound,
		}
	}

	if product.CustomerID != customer.ID {
		return &handler_err.InfoErr{
			Message: "this product does not exist in your cart",
			Err: handler_err.ErrNotFound,
		}
	}

	if err = dp.repository.DeleteProduct(cartInput.CartID); err != nil {
		return &handler_err.InfoErr{
			Message: "could not delete product",
			Err: handler_err.ErrInternal,
		}
	}

	return &handler_err.InfoErr{}
}