package handlers

import (
	"github.com/labstack/echo"
)

func RawResponse(c echo.Context, response interface{}, httpCode int) error {
	var responseFunc func(int, interface{}) error
	responseFunc = c.JSON
	return responseFunc(httpCode, response)
}
