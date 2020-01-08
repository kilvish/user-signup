package pkg

import (
	"net/http"
)

// Validate is like validator
func (request *CreateUserRequest) Validate() *Error {
	if request.Email == nil || *request.Email == "" {
		return &Error{http.StatusBadRequest, "Email id is missing"}
	}

	if request.Name == nil || *request.Name == "" {
		return &Error{http.StatusBadRequest, "Name  is missing"}
	}

	if request.Password == nil || *request.Password == "" {
		return &Error{http.StatusBadRequest, " Password is missing"}
	}

	if request.Request != nil {
		if typ, ok := interface{}(request.Request).(Validator); ok {
			if err := typ.Validate(); err != nil {
				return err
			}
		}
	}

	return nil
}
