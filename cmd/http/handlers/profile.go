package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kilvish/user-signup/cmd/http/pkg"
	"github.com/kilvish/user-signup/internal/core/users"
	"github.com/labstack/echo"
)

type ProfileHandler struct{}

func (handler ProfileHandler) Get(c echo.Context) error {
	var (
		r   pkg.UserProfileResponse
		err *pkg.Error
	)
	response := &r
	userResponse := new(pkg.ProfileResponse)
	respError := new(pkg.Error)

	requestID := c.Get("RequestID").(string)
	method := c.Get("Method").(string)
	req := new(pkg.GetUserRequest)

	if err = req.ExtractFromHTTP(c); err == nil {
		req.Request = &pkg.Request{RequestID: &requestID, Method: &method}
		err = req.Validate()
	}
	if err != nil {
		response.SetErrorData(&err.HttpCode, method, requestID)
		userResponse.Error = *err
		response.ResponseData = userResponse
		return RawResponse(c, response, err.HttpCode)
	}
	response, err = users.Profile(*req)
	if err != nil {
		respError.HttpCode = err.HttpCode
		respError.ErrMsg = err.ErrMsg
		userResponse.Error = *respError
		response.ResponseData = userResponse
		response.SetErrorData(&err.HttpCode, method, requestID)
		return RawResponse(c, response, err.HttpCode)
	}
	response.SetErrorData(nil, method, requestID)
	return RawResponse(c, response, http.StatusOK)
}

func (handler ProfileHandler) Put(c echo.Context) error {
	var (
		r   pkg.UpdatedProfileResponse
		err *pkg.Error
	)
	response := &r
	userResponse := new(pkg.ProfileUpdateResponse)
	respError := new(pkg.Error)

	requestID := c.Get("RequestID").(string)
	method := c.Get("Method").(string)
	req := new(pkg.ProfileUpdateRequest)

	if err = req.ExtractFromHTTP(c); err == nil {
		req.Request = &pkg.Request{RequestID: &requestID, Method: &method}
		s, _ := json.MarshalIndent(req, "", " ")
		fmt.Println(string(s))
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
	response, err = users.ProfileUpdate(*req)
	if err != nil {
		respError.HttpCode = err.HttpCode
		respError.ErrMsg = err.ErrMsg
		userResponse.Error = *respError
		response.ResponseData = userResponse
		response.SetErrorData(&err.HttpCode, method, requestID)
		return RawResponse(c, response, err.HttpCode)
	}
	response.SetErrorData(nil, method, requestID)
	return RawResponse(c, response, http.StatusOK)
}
