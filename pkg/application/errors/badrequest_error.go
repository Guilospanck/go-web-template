package errors

type BadRequestError struct {
	Message string
}

func (err BadRequestError) Error() string {
	return err.Message
}

func NewBadRequestError(message string) BadRequestError {
	return BadRequestError{
		Message: message,
	}
}
