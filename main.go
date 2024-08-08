package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mAmineChniti/FlixLib/pages"
	"github.com/mAmineChniti/FlixLib/utils"
	"os"
)

func main() {
	// Echo instance
	app := echo.New()
	// Routes
	app.GET("/", func(c echo.Context) error {
		templ.handler(utils.Render(pages.Index))
	})
	port := os.Getenv("PORT")
	// Start server
	app.Start(":" + port)
}
