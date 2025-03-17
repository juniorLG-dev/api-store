package usecase

import (
	"loja/internal/customer/application/domain"
	"loja/internal/common/smtp"
	"loja/internal/customer/adapter/output/repository"
	"loja/internal/customer/adapter/output/cache"
	"loja/internal/customer/adapter/output/model/customer_cache"
	"loja/internal/configuration/handler_err"
	"loja/internal/customer/application/dto"

	"strconv"
	"math/rand"
)

type RegisterCustomer struct {
	repository repository.PortRepository
	smtp smtp.PortSMTP
	cache cache.PortCache
}

func NewUseCaseRegisterCustomer(
	repo repository.PortRepository,
	smtp smtp.PortSMTP,
	cache cache.PortCache,
) *RegisterCustomer {
	return &RegisterCustomer{
		repository: repo,
		smtp: smtp,
		cache: cache,
	}
}

func generateCode() (codeString string) {
	for i := 0; i <= 6; i++ {
		codeString += strconv.Itoa(rand.Intn(10))
	}

	return
}

func (rc *RegisterCustomer) Run(customerInput dto.RegisterCustomerInput) *handler_err.InfoErr {
	code := generateCode()

	_, msgErr := domain.NewCustomerDomain(
		customerInput.Name,
		customerInput.Username,
		customerInput.Email,
		customerInput.Password,
	)
	if msgErr.Err != nil {
		return msgErr
	}

	if rc.repository.EmailExists(customerInput.Email) {
		return &handler_err.InfoErr{
			Message: "this email is already in use",
			Err: handler_err.ErrInvalidInput,
		}
	}

	if rc.repository.UsernameExists(customerInput.Username) {
		return &handler_err.InfoErr{
			Message: "this useraname is already in use",
			Err: handler_err.ErrInvalidInput,
		}
	}

	customerCache := customer_cache.NewInfoCustomer(
		customerInput.Name,
		customerInput.Username,
		customerInput.Email,
		customerInput.Password,
		code,
	)

	if err := rc.cache.SetCache(*customerCache); err != nil {
		return &handler_err.InfoErr{
			Message: "unable to set cache",
			Err: handler_err.ErrInternal,
		}
	}

	if err := rc.smtp.SendVerificationEmail(customerInput.Email, code); err != nil {
		return &handler_err.InfoErr{
			Message: "unable to send verification code",
			Err: handler_err.ErrInternal,
		}
	}
	
	return &handler_err.InfoErr{}
}