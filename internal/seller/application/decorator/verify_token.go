package decorator

import (
	"loja/internal/common/domain/service"
	"loja/internal/seller/application/usecase"
	"loja/internal/configuration/handler_err"
)

var token string

type TokenVerifier[V any, R any] struct {
	usecase usecase.Usecase[V, R]
	tokenGenerator service.TokenGenerator
}

func NewTokenVerifier[V any, R any](usecase usecase.Usecase[V, R]) usecase.Usecase[V, R] {
	return &TokenVerifier[V, R]{
		usecase: usecase,
		tokenGenerator: *service.NewTokenGenerator("seller"),
	}
}

func (tv *TokenVerifier[V, R]) Run(sellerInput V) (R, *handler_err.InfoErr) {
	_, msgErr := tv.tokenGenerator.VerifyToken(token)
	if msgErr.Err != nil {
		return *new(R), msgErr
	}

	output, msgErr := tv.usecase.Run(sellerInput)
	if msgErr.Err != nil {
		return *new(R), msgErr
	}

	return output, &handler_err.InfoErr{}
}

func SetToken(tokenValue string) {
	token = tokenValue
}