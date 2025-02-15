package usecase 

import (
	"loja/internal/seller/adapter/output/repository"
	"loja/internal/seller/adapter/output/cache"
	"loja/internal/seller/adapter/output/smtp"
	"loja/internal/seller/application/domain"
	"loja/internal/configuration/handler_err"
	"loja/internal/seller/application/dto"
	"loja/internal/common/domain/service"

	"fmt"
)

type CreateSeller struct {
	repository repository.PortRepository
	smtp smtp.PortSMTP
	cache cache.PortCache
}

func NewUseCaseCreateSeller(repo repository.PortRepository, smtp smtp.PortSMTP, cache cache.PortCache) *CreateSeller {
	return &CreateSeller{
		repository: repo,
		smtp: smtp,
		cache: cache,
	}
}

func (cs *CreateSeller) Run(sellerInput dto.CreateSellerInput) *handler_err.InfoErr {
	sellerCache, err := cs.cache.GetCache(sellerInput.Email)
	if err != nil {
		return &handler_err.InfoErr{
			Message: "expired code",
			Err: handler_err.ErrInvalidInput,
		}
	}

	verificationCode, msgErr := service.NewCode(sellerCache.Code)
	if msgErr.Err != nil {
		fmt.Println("error: ", msgErr)
		return msgErr
	}

	sellerDomain, msgErr := domain.NewSellerDomain(
		sellerCache.Name,
		sellerCache.Username,
		sellerCache.Email,
		sellerCache.Password,
	)

	if msgErr.Err != nil {
		return msgErr
	}

	if !verificationCode.CheckCode(sellerInput.Code) {
		return &handler_err.InfoErr{
			Message: "invalid code",
			Err: handler_err.ErrInvalidInput,
		}
	}

	if err := cs.repository.CreateSeller(*sellerDomain); err != nil {
		return &handler_err.InfoErr{
			Message: "unable to create user",
			Err: handler_err.ErrInternal,
		}
	}

	return &handler_err.InfoErr{}
}

