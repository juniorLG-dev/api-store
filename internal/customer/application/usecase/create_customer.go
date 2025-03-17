package usecase

import (
	"loja/internal/customer/adapter/output/cache"
	"loja/internal/customer/adapter/output/repository"
	"loja/internal/customer/application/dto"
	"loja/internal/customer/application/domain"
	"loja/internal/common/domain/service"
	"loja/internal/configuration/handler_err"
)

type CreateCustomer struct {
	repository repository.PortRepository
	cache cache.PortCache
}

func NewUseCaseCreateCustomer(
	repository repository.PortRepository,
	cache cache.PortCache,
) *CreateCustomer {
	return &CreateCustomer {
		repository: repository,
		cache: cache,
	}
}

func (cc *CreateCustomer) Run(customerInput dto.CreateCustomerInput) *handler_err.InfoErr {
	infoCustomerCache, err := cc.cache.GetCache(customerInput.Email)
	if err != nil {
		return &handler_err.InfoErr{
			Message: "expired code",
			Err: handler_err.ErrInvalidInput,
		}
	}

	code, msgErr := service.NewCode(infoCustomerCache.Code)
	if msgErr.Err != nil {
		return msgErr
	}

	if !code.CheckCode(customerInput.Code) {
		return &handler_err.InfoErr{
			Message: "invalid code",
			Err: handler_err.ErrInvalidInput,
		}
	}

	customerDomain, msgErr := domain.NewCustomerDomain(
		infoCustomerCache.Name,
		infoCustomerCache.Username,
		infoCustomerCache.Email,
		infoCustomerCache.Password,
	)
	if msgErr.Err != nil {
		return msgErr
	}

	if err := cc.repository.CreateCustomer(*customerDomain); err != nil {
		return &handler_err.InfoErr{
			Message: "unable to create customer",
			Err: handler_err.ErrInternal,
		}
	}

	return &handler_err.InfoErr{}
}
