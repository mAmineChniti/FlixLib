package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	port      = ":3500"
	stylesDir = "styles"
)

var (
	logger = log.New(os.Stdout, "HTTP Server: ", log.LstdFlags)
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "Request: Method=${method}, Uri=${uri}, Status=${status}\n",
	}))
	e.Use(middleware.CORS())

	// Routes
	e.GET("/", pageServer)
	e.GET("/status", statusCheck)

	//Route for Serving Static Components
	e.GET("/components/:Component", func(c echo.Context) error {
		return componentServer(c, nil)
	})

	// Static Routes
	e.Static("/styles", stylesDir)

	// Start server
	logger.Printf("Server started on port %s\n", port)
	e.Logger.Fatal(e.Start(port))
}
