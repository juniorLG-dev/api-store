package vo

import (
	"loja/internal/configuration/handler_err"
)

type Price struct {
	Value float64
}

func NewPrice(value float64) (*Price, *handler_err.InfoErr) {
	if value < 2.50 {
		return &Price{}, &handler_err.InfoErr{
			Message: "its price must be greater than 2.50",
			Err: handler_err.ErrInvalidInput,
		}
	}

	return &Price{
		Value: value,
	}, &handler_err.InfoErr{}
}