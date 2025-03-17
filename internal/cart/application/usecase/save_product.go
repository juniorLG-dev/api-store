package usecase

import (
	"loja/internal/configuration/handler_err"
	"loja/internal/cart/adapter/output/repository"
	"loja/internal/cart/adapter/output/gateway"
	"loja/internal/cart/application/dto"
	"loja/internal/cart/application/domain/entities"
	"loja/internal/common/domain/service"

)

type SaveProduct struct {
	repository repository.PortRepository
	tokenGenerator *service.TokenGenerator
	gateway gateway.PortGateway
}

func NewUseCaseSaveProduct(
	repository repository.PortRepository,
	gateway gateway.PortGateway,
) *SaveProduct {
	tokenGenerator := service.NewTokenGenerator("customer")
	return &SaveProduct{
		repository: repository,
		tokenGenerator: tokenGenerator,
		gateway: gateway,
	}
}

func (sp *SaveProduct) Run(cartInput dto.SaveProductInput, token string) *handler_err.InfoErr {
	customer, msgErr := sp.tokenGenerator.VerifyToken(token)
	if msgErr.Err != nil {
		return msgErr
	}

	if customer.Type != sp.tokenGenerator.Value{
		return &handler_err.InfoErr{
			Message: "you do not have permission to save a product",
			Err: handler_err.ErrUnauthorized,
		}
	}

	product, err := sp.gateway.GetProductByID(cartInput.ProductID)
	if err != nil {
		return &handler_err.InfoErr{
			Message: "product not found",
			Err: handler_err.ErrNotFound,
		}
	}

	cartDomain, msgErr := entities.NewCartDomain(
		product.Description,
		product.Price,
		cartInput.ProductID,
		customer.ID,
		product.SellerID,
	)
	if msgErr.Err != nil {
		return msgErr
	}

	if err := sp.repository.SaveProduct(*cartDomain); err != nil {
		return &handler_err.InfoErr{
			Message: "could not save product",
			Err: handler_err.ErrInternal,
		}
	}

	return &handler_err.InfoErr{}
}