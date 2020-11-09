package exceptions

import "fmt"

//UnsupportedValueError represents a validation error when value is not supported for algorithm
type UnsupportedValueError struct {
	Reason string
}

func (u *UnsupportedValueError) Error() string {
	return fmt.Sprintf(u.Reason)
}
