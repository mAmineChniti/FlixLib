package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
)

func main() {
	// Constants
	pagesDir := "pages/"
	componentsDir := "components/"
	port := ":3500"
	stylesDir := "styles/"
	logger := log.New(os.Stdout, "HTTP Server: ", log.LstdFlags)
	//Flogger := log.New(os.File.Write(), "HTTP Server: ", log.LstdFlags)
	e := echo.New()

	e.GET("/status", func(c echo.Context) error {
		logger.Printf("Working Fine!\n")
		return c.String(http.StatusOK, "Working Fine!\n")
	})

	e.GET("/", func(c echo.Context) error {
		urlPath := c.Request().URL.Path
		logger.Printf("Request: %s\n", urlPath)

		if urlPath == "/" {
			urlPath = "/index"
		}

		filePath := filepath.Join(pagesDir, urlPath+".html")
		_, err := os.Stat(filePath)
		if err != nil {
			logger.Printf("File not found: %s\n", filePath)
			return c.String(http.StatusNotFound, "File not found\n")
		}

		logger.Printf("Serving file: %s\n", filePath)

		tmpl := template.Must(template.ParseFiles(filePath))
		return tmpl.Execute(c.Response(), nil)
	})

	e.GET("/api/:Componet", func(c echo.Context) error {
		componentName := strings.ToLower(c.Param("Componet"))
		filePath := filepath.Join(componentsDir, componentName+".html")
		_, err := os.Stat(filePath)
		if err != nil {
			logger.Printf("API component file not found: %s\n", filePath)
			return c.String(http.StatusNotFound, "API component file not found\n")
		}

		logger.Printf("Serving API component: %s\n", filePath)

		tmpl := template.Must(template.ParseFiles(filePath))
		return tmpl.Execute(c.Response(), nil)
	})
	e.Static("/styles", stylesDir)
	logger.Printf("Server started on port %s\n", port)
	e.Logger.Fatal(e.Start(port))
}
