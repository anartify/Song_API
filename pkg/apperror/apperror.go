package apperror

// This struct helps in writing custom error messages
type CustomError struct {
	Message string
}

// Error() method returns the string containing the error message
func (e *CustomError) Error() string {
	return e.Message
}
