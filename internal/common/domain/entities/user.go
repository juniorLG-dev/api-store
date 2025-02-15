package entities

import (
	"loja/internal/common/domain/vo"
	"loja/internal/configuration/handler_err"
)

type User struct {
	ID vo.ID
	Name string
	Username vo.Username
}

func NewUser(id vo.ID, name, username string) (*User, *handler_err.InfoErr) {
	usernameVerify, msgErr := vo.NewUsername(username)
	if msgErr.Err != nil {
		return &User{}, msgErr
	}

	return &User{
		ID: id,
		Name: name,
		Username: *usernameVerify,
	}, &handler_err.InfoErr{}
}

func (u *User) GetID() string {
	return u.ID.Value
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetUsername() string {
	return u.Username.Value
}