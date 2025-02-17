package decorator

import (
	"loja/internal/common/domain/service"
	"loja/internal/configuration/handler_err"
	"loja/internal/seller/application/query"
)

type TokenVerifierInput[V any] struct {
	Token string
	Data V
}

type TokenVerifier[V any, R any] struct {
	query query.Query[V, R]
	tokenGenerator service.TokenGenerator
}

func NewTokenVerifier[V any, R any](query query.Query[V, R]) query.Query[TokenVerifierInput[V], R] {
	return &TokenVerifier[V, R]{
		query: query,
		tokenGenerator: *service.NewTokenGenerator("seller"),
	}
}


func (tv *TokenVerifier[V, R]) Run(input TokenVerifierInput[V]) (R, *handler_err.InfoErr) {
	_, msgErr := tv.tokenGenerator.VerifyToken(input.Token)
	if msgErr.Err != nil {
		return *new(R), msgErr
	}

	output, msgErr := tv.query.Run(input.Data)
	if msgErr.Err != nil {
		return *new(R), msgErr
	}

	return output, &handler_err.InfoErr{}
}
