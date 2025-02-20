package vo

import (
	"loja/internal/configuration/handler_err"
)

type Price struct {
	Value float64
}

func NewPrice(price float64) (*Price, *handler_err.InfoErr) {
	if price < 2.50 {
		return &Price{}, &handler_err.InfoErr{
			Message: "its price must be greater than 2.50",
			Err: handler_err.ErrInvalidInput,
		}
	}

	return &Price{
		Value: price,
	}, &handler_err.InfoErr{}
}