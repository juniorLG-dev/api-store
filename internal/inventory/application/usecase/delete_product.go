package usecase

import (
	"loja/internal/inventory/adapter/output/repository"
	"loja/internal/inventory/application/dto"
	"loja/internal/configuration/handler_err"
	"loja/internal/common/domain/service"
)

type DeleteProduct struct {
	repository repository.PortRepository
	tokenGenerator *service.TokenGenerator
}

func NewUseCaseDeleteProduct(repo repository.PortRepository) *DeleteProduct {
	tokenGenerator := service.NewTokenGenerator("seller")
	return &DeleteProduct {
		repository: repo,
		tokenGenerator: tokenGenerator,
	}
}

func (dp *DeleteProduct) Run(productInput dto.DeleteProductInput, token string) *handler_err.InfoErr {
	seller, msgErr := dp.tokenGenerator.VerifyToken(token)
	if msgErr.Err != nil {
		return msgErr
	}

	if seller.Type != dp.tokenGenerator.Value {
		return &handler_err.InfoErr{
			Message: "you do not have permission to perform this action",
			Err: handler_err.ErrUnauthorized,
		}
	}

	product, err := dp.repository.GetProductByID(productInput.ID)
	if err != nil {
		return &handler_err.InfoErr{
			Message: "product not found",
			Err: handler_err.ErrNotFound,
		}
	}
	

	if seller.ID != product.SellerID {
		return &handler_err.InfoErr{
			Message: "you do not have permission to delete this product",
			Err: handler_err.ErrUnauthorized,
		}
	}

	if err := dp.repository.DeleteProduct(productInput.ID); err != nil {
		return &handler_err.InfoErr{
			Message: "could not delete product",
			Err: handler_err.ErrInternal,
		}
	}

	return &handler_err.InfoErr{}
}