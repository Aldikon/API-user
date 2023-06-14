package model

type ErrorValid struct {
	Message string
}

func (e ErrorValid) Error() string {
	return "validation" + e.Message
}

func NewValidationError(m string) *ErrorValid {
	return &ErrorValid{
		Message: m,
	}
}
