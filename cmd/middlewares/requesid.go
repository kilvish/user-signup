package middlewares

import (
	"github.com/kilvish/user-signup/cmd/http/pkg/uuid"
	"github.com/labstack/echo"
)

// RequestID middleware for setting request id
func RequestID(next echo.HandlerFunc) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		c.Set("RequestID", uuid.GetUUID())
		return next(c)
	})
}
