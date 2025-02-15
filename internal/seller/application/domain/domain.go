package domain

import (
	"loja/internal/common/domain/vo"
	"loja/internal/configuration/handler_err"
)

type SellerDomain struct {
	ID vo.ID
	Name string
	Username vo.Username
	Email string
	Password vo.Password
}

func NewSellerDomain(
	name string,
	username string,
	email string,
	pass  string,
) (*SellerDomain, *handler_err.InfoErr) {
	usernameVerify, msgErr := vo.NewUsername(username)
	if msgErr.Err != nil {
		return &SellerDomain{}, msgErr
	}

	return &SellerDomain{
		ID: *vo.NewID(),
		Name: name,
		Username: *usernameVerify,
		Email: email,
		Password: *vo.NewPassword(pass),
	}, &handler_err.InfoErr{}
}


