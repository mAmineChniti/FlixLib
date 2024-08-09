package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mAmineChniti/FlixLib/pages"
	"github.com/mAmineChniti/FlixLib/utils"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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
	app.GET("/component/:componentName", func(c echo.Context) error {
		componentName := c.Param("componentName")
		return utils.RenderComponent(c, componentName)
	})
	app.HTTPErrorHandler = func(err error, c echo.Context) {
		if c.Response().Committed {
			return
		}
		if he, ok := err.(*echo.HTTPError); ok && he.Code == http.StatusNotFound {
			_ = utils.Render(c, pages.NotFound())
		} else {
			c.Echo().DefaultHTTPErrorHandler(err, c)
		}
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server in a goroutine
	go func() {
		logger.Println("Starting server on port " + port)
		if err := app.Start(":" + port); err != nil && err != http.ErrServerClosed {
			app.Logger.Fatal("Shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	logger.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := app.Shutdown(ctx); err != nil {
		app.Logger.Fatal("Server forced to shutdown:", err)
	}

	logger.Println("Server exiting")
}
