package handlers

import (
	"net/http"

	"github.com/kilvish/user-signup/cmd/http/pkg"
	"github.com/kilvish/user-signup/internal/core/users"
	"github.com/labstack/echo"
)

// SignUpHandler handler
type SignUpHandler struct{}

//Create creates an user
func (handler SignUpHandler) Create(c echo.Context) error {
	var (
		r   pkg.UserSignUPResponse
		err *pkg.Error
	)
	response := &r
	userResponse := new(pkg.UserResponse)
	respError := new(pkg.Error)
	requestID := c.Get("RequestID").(string)
	method := c.Get("Method").(string)
	req := new(pkg.CreateUserRequest)
	if err = req.ExtractFromHTTP(c); err == nil {
		req.Request = &pkg.Request{RequestID: &requestID, Method: &method}
		err = req.Validate()
	}
	if err != nil {
		response.SetErrorData(&err.HttpCode, method, requestID)
		respError.HttpCode = err.HttpCode
		respError.ErrMsg = err.ErrMsg
		userResponse.Error = *respError
		response.ResponseData = userResponse
		return RawResponse(c, response, err.HttpCode)
	}
	response, err = users.SignUP(*req)
	if err != nil {
		respError.HttpCode = err.HttpCode
		respError.ErrMsg = err.ErrMsg
		userResponse.Error = *respError
		response.SetErrorData(&err.HttpCode, method, requestID)
		response.ResponseData = userResponse
		return RawResponse(c, response, err.HttpCode)
	}
	response.SetErrorData(nil, method, requestID)
	return RawResponse(c, response, http.StatusOK)
}
