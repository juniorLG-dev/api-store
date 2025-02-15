package handler_err

import (
	"errors"
	"net/http"
)

var (
	ErrInternal = errors.New("internal error")
	ErrInvalidInput = errors.New("invalid input error")
	ErrNotFound = errors.New("not found error")
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
}

func HandlerErr(infoErr *InfoErr) *HandlerError {
	return errorsHTTP[infoErr.Err](infoErr.Message, infoErr.Err)
}