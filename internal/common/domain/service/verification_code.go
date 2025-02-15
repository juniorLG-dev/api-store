package service

import (
	"loja/internal/configuration/handler_err"
	"regexp"
)

type CodeGenerator struct {
	Value string
}

func NewCode(code string) (*CodeGenerator, *handler_err.InfoErr) {
	match, _ := regexp.MatchString(`^\d{7}$`, code)
	if !match {
		return &CodeGenerator{}, &handler_err.InfoErr{
			Message: "your code must contain only numbers (0-9) and must be 7 digits long",
			Err: handler_err.ErrInvalidInput,
		}
	}

	return &CodeGenerator{
		Value: code,
	}, &handler_err.InfoErr{}
}

func (c *CodeGenerator) CheckCode(code string) bool {
	return c.Value == code
}