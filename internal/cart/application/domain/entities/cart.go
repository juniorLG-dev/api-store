package entities

import (
	"loja/internal/configuration/handler_err"
	"loja/internal/cart/application/domain/vo"
)

type CartDomain struct {
	ID 					vo.ID
	Description string
	Price       vo.Price
	ProductID   string
	CustomerID  string
	SellerID    string
}

func NewCartDomain(
	description string,
	price float64,
	productID string,
	customerID string,
	sellerID string,
) (*CartDomain, *handler_err.InfoErr) {
	priceValue, msgErr := vo.NewPrice(price)
	if msgErr.Err != nil {
		return &CartDomain{}, msgErr
	}

	return &CartDomain{
		ID: *vo.NewID(),
		Description: description,
		Price: *priceValue,
		ProductID: productID,
		CustomerID: customerID,
		SellerID: sellerID,
	}, &handler_err.InfoErr{}
}