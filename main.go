package main
import (
	"os"
"github.com/labstack/echo/v4"
)
func main() {
// Echo instance
	app := echo.New()
// Routes
	// get page names from pages/ directory then turn those page names into our routes
	for _, page := range pages {
	app.GET("/"+page, func(c echo.Context) error {
		return c.File("pages/" + c.Param("page") + ".html")
	})
	}
	port := os.Getenv("PORT")
	// Start server
	app.Start(":" + port)
}
