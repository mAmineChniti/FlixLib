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

// Constants
const (
	pagesDir      = "pages/"
	componentsDir = "components/"
	port          = ":3500"
	stylesDir     = "styles/"
)

var (
	logger = log.New(os.Stdout, "HTTP Server: ", log.LstdFlags)
)

func main() {
	e := echo.New()

	// Routes
	e.GET("/", pageServer)
	e.GET("/status", statusCheck)

	// Static Routes
	e.Static("/styles", stylesDir)

	// Start server
	logger.Printf("Server started on port %s\n", port)
	e.Logger.Fatal(e.Start(port))
}

func statusCheck(c echo.Context) error {
	logger.Println("Working Fine!")
	return c.String(http.StatusOK, "Working Fine!\n")
}

func pageServer(c echo.Context) error {
	urlPath := c.Request().URL.Path
	logger.Printf("Request: %s\n", urlPath)

	if urlPath == "/" {
		urlPath = "/index"
	}

	filePath := filepath.Join(pagesDir, urlPath+".html")
	if _, err := os.Stat(filePath); err != nil {
		logger.Printf("File not found: %s\n", filePath)
		return c.String(http.StatusNotFound, "File not found\n")
	}

	logger.Printf("Serving file: %s\n", filePath)

	tmpl := template.Must(template.ParseFiles(filePath))
	return tmpl.Execute(c.Response(), nil)
}

/*func componentServer(c echo.Context, data any) error {
	componentName := strings.ToLower(c.Param("Component"))
	filePath := filepath.Join(componentsDir, componentName+".html")
	if _, err := os.Stat(filePath); err != nil {
		logger.Printf("API component file not found: %s\n", filePath)
		return c.String(http.StatusNotFound, "API component file not found\n")
	}

	logger.Printf("Serving API component: %s\n", filePath)

	tmpl := template.Must(template.ParseFiles(filePath))
	return tmpl.Execute(c.Response(), data)
}
*/
