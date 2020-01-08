package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// ExtractFromHTTP reads http request body
func (receiver *CreateUserRequest) ExtractFromHTTP(c echo.Context) *Error {
	request := c.Request()
	if request.ContentLength > 0 {
		if err := json.NewDecoder(request.Body).Decode(receiver); err != nil {
			fmt.Println(err)
			return &Error{http.StatusBadRequest, err.Error()}
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

func (receiver *ProfileUpdateRequest) ExtractFromHTTP(c echo.Context) *Error {
	request := c.Request()
	if request.ContentLength > 0 {
		if err := json.NewDecoder(request.Body).Decode(receiver); err != nil {
			fmt.Println(err)
			return &Error{http.StatusBadRequest, err.Error()}
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
