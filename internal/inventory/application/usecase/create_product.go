package usecase

import (
	"loja/internal/inventory/adapter/output/repository"
	"loja/internal/inventory/application/domain/entities"
	"loja/internal/inventory/application/dto"
	"loja/internal/common/domain/service"
	"loja/internal/configuration/handler_err"
)

type CreateProduct struct {
	repository repository.PortRepository
	tokenGenerator *service.TokenGenerator
}

func NewUseCaseCreateProduct(repo repository.PortRepository) *CreateProduct {
	tokenGenerator := service.NewTokenGenerator("seller")
	return &CreateProduct{
		repository: repo,
		tokenGenerator: tokenGenerator,
	}
}

func (cp *CreateProduct) Run(productInput dto.CreateProductInput, token string) *handler_err.InfoErr {
	seller, msgErr := cp.tokenGenerator.VerifyToken(token)
	if msgErr.Err != nil {
		return msgErr
	}

	if seller.Type != cp.tokenGenerator.Value {
		return &handler_err.InfoErr{
			Message: "you do not have permission to create a product",
			Err: handler_err.ErrUnauthorized,
		}
	}

	productInventory, msgErr := entities.NewProductInventory(
		productInput.Description,
		productInput.Price,
		productInput.Quantity,
		seller.ID,
	)

	if msgErr.Err != nil {
		return msgErr
	}

	if err := cp.repository.CreateProduct(*productInventory); err != nil {
		return &handler_err.InfoErr{
			Message: "unable to create product",
			Err: handler_err.ErrUnauthorized,
		}
	}

	return &handler_err.InfoErr{}
}