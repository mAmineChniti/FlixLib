package main

import (
	"net/http"
	"log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mAmineChniti/FlixLib/pages"
	"github.com/mAmineChniti/FlixLib/utils"
	"os"
)

func main() {
	// Echo instance
	app := echo.New()
	// Middleware & Logger
	var logger = log.New(os.Stdout, "HTTP Server: ", log.LstdFlags)
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "Request: Method=${method}, Uri=${uri}, Status=${status}\n",
	}))
	app.Use(middleware.CORS())
	// Routes
	app.GET("/status", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	app.GET("/", func(c echo.Context) error {
		return utils.Render(c, pages.Index())
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// Start server
	logger.Println("Starting server on port " + port)
	app.Logger.Fatal(app.Start(":" + port))
}
