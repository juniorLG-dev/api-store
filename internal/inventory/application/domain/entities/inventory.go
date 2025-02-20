package entities

import (
	"loja/internal/configuration/handler_err"
	"loja/internal/inventory/application/domain/vo"
)

type ProductInventory struct {
	ID vo.ID
	Description string
	Price vo.Price
	Quantity vo.Quantity
	SellerID string
}

func NewProductInventory(
	description string,
	price float64,
	quantity int,
	sellerID string,
) (*ProductInventory, *handler_err.InfoErr) {
	priceValue, msgErr := vo.NewPrice(price)
	if msgErr.Err != nil {
		return &ProductInventory{}, msgErr
	}

	quantityValue, msgErr := vo.NewQuantity(quantity)
	if msgErr.Err != nil {
		return &ProductInventory{}, msgErr
	}

	return &ProductInventory{
		ID: *vo.NewID(),
		Description: description,
		Price: *priceValue,
		Quantity: *quantityValue,
		SellerID: sellerID,
	}, &handler_err.InfoErr{}
}

