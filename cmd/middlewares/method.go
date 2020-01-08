package middlewares

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

// Method middleware for setting method
func Method(next echo.HandlerFunc) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		method := c.Request().Method
		customMethod := strings.ToUpper(c.QueryParam("_method"))
		switch customMethod {
		case http.MethodGet, http.MethodDelete, http.MethodPost, http.MethodPut:
			method = customMethod
		}
		c.Request().Method = method
		c.SetRequest(c.Request())
		c.Set("Method", method)
		return next(c)
	})
}
