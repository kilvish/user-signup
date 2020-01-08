package pkg

import (
	"errors"
	"net/http"
)

//BaseResponse baseresponse
type BaseResponse struct {
	RequestID string `json:"request_id"`
	Method    string `json:"method"`
	HTTPCode  int    `json:"http_code"`
}

type Error struct {
	HttpCode int    `json:"code"`
	ErrMsg   string `json:""error`
}

var (
	ErrNotFound         = errors.New("Not Found")
	ErrResourceConflict = errors.New("Duplicate resource")
)

func (response *UserSignUPResponse) SetErrorData(httpCode *int, method, requestID string) {
	if httpCode != nil {
		response.BaseResponse.HTTPCode = *httpCode
	} else {
		response.BaseResponse.HTTPCode = http.StatusOK
	}
	response.BaseResponse.Method = method
	response.BaseResponse.RequestID = requestID
}

func (response *UserProfileResponse) SetErrorData(httpCode *int, method, requestID string) {
	if httpCode != nil {
		response.BaseResponse.HTTPCode = *httpCode
	} else {
		response.BaseResponse.HTTPCode = http.StatusOK
	}
	response.BaseResponse.Method = method
	response.BaseResponse.RequestID = requestID
}

func (response *UpdatedProfileResponse) SetErrorData(httpCode *int, method, requestID string) {
	if httpCode != nil {
		response.BaseResponse.HTTPCode = *httpCode
	} else {
		response.BaseResponse.HTTPCode = http.StatusOK
	}
	response.BaseResponse.Method = method
	response.BaseResponse.RequestID = requestID
}
