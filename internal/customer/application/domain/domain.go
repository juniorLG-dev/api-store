package domain

import (
	"loja/internal/common/domain/vo"
	"loja/internal/configuration/handler_err"
)

type CustomerDomain struct {
	ID vo.ID
	Name string
	Username vo.Username
	Email string
	Password vo.Password
}

func NewCustomerDomain(
	name string,
	username string,
	email string,
	password string,
) (*CustomerDomain, *handler_err.InfoErr) {
	usernameVerified, msgErr := vo.NewUsername(username)
	if msgErr.Err != nil {
		return &CustomerDomain{}, msgErr
	}

	return &CustomerDomain{
		ID: *vo.NewID(),
		Name: name,
		Username: *usernameVerified,
		Email: email,
		Password: *vo.NewPassword(password),
	}, &handler_err.InfoErr{}
}