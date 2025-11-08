package habit

import "fmt"

// InvalidInputError is returned when the input data is invalid
type InvalidInputError struct {
	field  string
	reason string
}

// Error implements error
func (i InvalidInputError) Error() string {
	return fmt.Sprintf("invalid input in field %s : %s", i.field, i.reason)
}
