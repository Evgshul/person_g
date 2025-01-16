package errors

import "errors"

var (
	ErrPersonNotFound = errors.New("Person not found")
	ErrInvalidPerson  = errors.New("Invalid person data")
)
