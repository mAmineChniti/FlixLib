package utils

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/mAmineChniti/FlixLib/components"
	"net/http"
)

var comps = map[string]templ.Component{
	"box": components.Box(),
}

func Render(c echo.Context, comp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	if err := comp.Render(c.Request().Context(), c.Response().Writer); err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return nil
}

func RenderComponent(c echo.Context, compName string) error {
	comp, exists := comps[compName]
	if !exists {
		return c.String(http.StatusNotFound, "Not Found")
	}
	return Render(c, comp)
}
