package query

import (
	"loja/internal/configuration/handler_err"
)

type Query[V any, R any] interface {
	Run(V) (R, *handler_err.InfoErr)
}