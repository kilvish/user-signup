package pkg

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
)

// ExtractFromHTTP reads http request body
func (receiver *GetUserRequest) ExtractFromHTTP(c echo.Context) *Error {
	request := c.Request()
	if request.ContentLength > 0 {
		if err := json.NewDecoder(request.Body).Decode(receiver); err != nil {
			return &Error{http.StatusBadRequest, "request format is invalid"}
		}
	}

	if receiver.Password == nil {
		Password := c.QueryParam("password")
		if Password != "" {
			receiver.Password = &Password
		}
	}

	if receiver.Email == nil {
		Email := c.QueryParam("email")
		if Email != "" {
			receiver.Email = &Email
		}
	}
	tmp := new(Request)
	if newrequest, ok := interface{}(receiver.Request).(Extractor); ok {
		receiver.Request = new(Request)
		if err := newrequest.ExtractFromHTTP(c); err != nil {
			return err
		}
		if *receiver.Request == *tmp {
			receiver.Request = nil
		}
	}
	return nil
}
