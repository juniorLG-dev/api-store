package usecase

import (
	"loja/internal/seller/application/domain"
	"loja/internal/seller/adapter/output/cache"
	"loja/internal/seller/adapter/output/model/seller_cache"
	"loja/internal/seller/adapter/output/repository"
	"loja/internal/seller/adapter/output/smtp"
	"loja/internal/configuration/handler_err"
	"loja/internal/seller/application/dto"

	"strconv"
	"math/rand"
)

type RegisterSeller struct {
	repository repository.PortRepository
	smtp smtp.PortSMTP
	cache cache.PortCache
}

func NewUseCaseRegisterSeller(
	repository repository.PortRepository,
	smtp smtp.PortSMTP,
	cache cache.PortCache,
) *RegisterSeller {
	return &RegisterSeller{
		repository: repository,
		smtp: smtp,
		cache: cache,
	}
}

func generateCode() string {
	var code string
	for i := 0; i <= 6; i++ {
		code += strconv.Itoa(rand.Intn(10))
	}

	return code
}

func (rs *RegisterSeller) Run(sellerInput dto.RegisterSellerInput) *handler_err.InfoErr {
	code := generateCode()

	_, msgErr := domain.NewSellerDomain(
		sellerInput.Name,
		sellerInput.Username,
		sellerInput.Email,
		sellerInput.Password,
	)
	if msgErr.Err != nil {
		return msgErr
	}

	if rs.repository.UsernameExists(sellerInput.Username) {
		return &handler_err.InfoErr{
			Message: "this username is already in use",
			Err: handler_err.ErrInvalidInput,
		}
	}

	if rs.repository.EmailExists(sellerInput.Email) {
		return &handler_err.InfoErr{
			Message: "this email in already in use",
			Err: handler_err.ErrInvalidInput,
		}
	}

	sellerCache := seller_cache.NewInfoSeller(
		sellerInput.Name,
		sellerInput.Username,
		sellerInput.Email,
		sellerInput.Password,
		code,
	)
	
	if err := rs.cache.SetCache(*sellerCache); err != nil {
		return &handler_err.InfoErr{
			Message: "unable to register seller",
			Err: handler_err.ErrInternal,
		}
	}

	if err := rs.smtp.SendVerificationEmail(sellerInput.Email, code); err != nil {
		return &handler_err.InfoErr{
			Message: "unable to send verification code",
			Err: handler_err.ErrInternal,
		}
	}

	return &handler_err.InfoErr{}
}