package apperror

import "fmt"

// This struct helps in writing custom error messages
type CustomError struct {
	Message string
}

// Error() method returns the string containing the error message
func (e *CustomError) Error() string {
	return e.Message
}

// Combine([]error) method combines multiple errors into one
func (e *CustomError) Combine(errors []error) error {
	var combinedErr error
	for _, err := range errors {
		if combinedErr == nil {
			combinedErr = err
		} else {
			combinedErr = fmt.Errorf("%v; %v", combinedErr, err)
		}
	}
	e.Message = combinedErr.Error()
	return combinedErr
}
