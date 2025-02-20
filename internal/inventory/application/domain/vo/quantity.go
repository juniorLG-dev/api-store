package vo

import (
	"loja/internal/configuration/handler_err"
)

type Quantity struct {
	Value int
}

func NewQuantity(quantity int) (*Quantity, *handler_err.InfoErr) {
	if quantity <= 0 {
		return &Quantity{}, &handler_err.InfoErr{
			Message: "its quantity must be greater than 0",
			Err: handler_err.ErrInvalidInput,
		}
	}

	return &Quantity{
		Value: quantity,
	}, &handler_err.InfoErr{}
}