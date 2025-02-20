package handler_err

import (
	"errors"
	"net/http"
)

var (
	ErrInternal = errors.New("internal error")
	ErrInvalidInput = errors.New("invalid input error")
	ErrNotFound = errors.New("not found error")
	ErrUnauthorized = errors.New("Unauthorized error")
)

type InfoErr struct {
	Message string
	Err error
}

type HandlerError struct {
	Message string
	Err string
	Status int
}

func (i *InfoErr) Internal() *HandlerError {
	return &HandlerError{
		Message: i.Message,
		Err: i.Err.Error(),
		Status: http.StatusInternalServerError,
	}
}

func (i *InfoErr) InvalidInput() *HandlerError {
	return &HandlerError{
		Message: i.Message,
		Err: i.Err.Error(),
		Status: http.StatusBadRequest,
	}
}

func (i *InfoErr) NotFound() *HandlerError {
	return &HandlerError{
		Message: i.Message,
		Err: i.Err.Error(),
		Status: http.StatusNotFound,
	}
}

func (i *InfoErr) Unauthorized() *HandlerError {
	return &HandlerError{
		Message: i.Message,
		Err: i.Err.Error(),
		Status: http.StatusUnauthorized,
	}
}

var errorsHTTP = map[error]func(string, error)*HandlerError{
	ErrInternal: func(msg string, err error) *HandlerError {
		return (&InfoErr{
			Message: msg,
			Err: err,
		}).Internal()
	},
	ErrInvalidInput: func(msg string, err error) *HandlerError {
		return (&InfoErr{
			Message: msg,
			Err: err,
		}).InvalidInput()
	},
	ErrNotFound: func(msg string, err error) *HandlerError {
		return (&InfoErr{
			Message: msg,
			Err: err,
		}).NotFound()
	},
	ErrUnauthorized: func(msg string, err error) *HandlerError {
		return (&InfoErr{
			Message: msg,
			Err: err,
		}).Unauthorized()
	},
}

func HandlerErr(infoErr *InfoErr) *HandlerError {
	return errorsHTTP[infoErr.Err](infoErr.Message, infoErr.Err)
}