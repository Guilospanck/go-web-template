package errors

type UnauthorizedError struct {
	Message string
}

func (err UnauthorizedError) Error() string {
	return err.Message
}

func NewUnauthorizedError(message string) UnauthorizedError {
	return UnauthorizedError{
		Message: message,
	}
}
