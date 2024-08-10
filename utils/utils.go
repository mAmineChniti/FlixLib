package utils

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Render(c echo.Context, comp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	if err := comp.Render(c.Request().Context(), c.Response().Writer); err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return nil
}
