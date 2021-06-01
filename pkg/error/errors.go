package error

import (
	"errors"
	"fmt"
)

//Reference article:
//https://hackernoon.com/golang-handling-errors-gracefully-8e27f1db729f

// CustomError struct
type CustomError struct {
	errType      ErrorType
	wrappedError error
}

// Error return error message
func (err CustomError) Error() string {
	return err.wrappedError.Error()
}

// Stacktrace return error stacktrace message
func (err CustomError) Stacktrace() string {
	return fmt.Sprintf("%+v\n", err.wrappedError)
}

// New creates a no type error
func New(msg string) error {
	return CustomError{errType: Error, wrappedError: errors.New(msg)}
}
