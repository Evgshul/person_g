package errors

type NotFound struct {
	Message string
}

type InvalidArgument struct {
	Message string
}

func (e NotFound) Error() string {
	return e.Message
}

func (e InvalidArgument) Error() string {
	return e.Message
}
