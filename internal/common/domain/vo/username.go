package vo

import (
	"loja/internal/configuration/handler_err"
	"regexp"
)

type Username struct {
	Value string
}

func NewUsername(username string) (*Username, *handler_err.InfoErr) {
	match, _ := regexp.MatchString("^[a-zA-Z0-9_]+$", username)
	if !match {
		return &Username{}, &handler_err.InfoErr{
			Message: "your username must contain only uppercase and/or lowercase letters, 0-9 and \"_\"",
			Err: handler_err.ErrInvalidInput,
		}
	}

	return &Username{
		Value: username,
	}, &handler_err.InfoErr{}
}