package models

type Error struct {
	Message    string   `json:"message"`
	Code       int      `json:"code"`
	Name       string   `json:"name"`
	Error      error    `json:"-"`
	Validation []string `json:"validation,omitempty"`
}

func BindError() *Error {
	return &Error{Code: 400, Message: "Error processing request.", Name: "BIND_ERROR"}
}

func ValidationError(errors []string) *Error {
	return &Error{Code: 400, Name: "VALIDATION", Message: "A validation error occurred.", Validation: errors}
}
