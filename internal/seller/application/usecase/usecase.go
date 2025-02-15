package usecase

import (
	"loja/internal/configuration/handler_err"
)

type Usecase[V any, R any] interface {
	Run(V) (R, *handler_err.InfoErr)
}