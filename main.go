package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	pagesDir := "pages/"
	componentsDir := "components/"
	port := 3500

	// Define a custom logger with a prefix
	logger := log.New(os.Stdout, "HTTP Server: ", log.LstdFlags)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		urlPath := r.URL.Path
		logger.Printf("Request: %s\n", urlPath)

		if urlPath == "/" {
			urlPath = "/index"
		}

		filePath := filepath.Join(pagesDir, urlPath+".html")
		_, err := os.Stat(filePath)
		if err != nil {
			logger.Printf("File not found: %s\n", filePath)
			http.NotFound(w, r)
			return
		}

		logger.Printf("Serving file: %s\n", filePath)

		tmpl := template.Must(template.ParseFiles(filePath))
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/api/Box", func(w http.ResponseWriter, r *http.Request) {
		componentName := r.URL.Path[len("/api/"):]
		filePath := filepath.Join(componentsDir, componentName+".html")
		_, err := os.Stat(filePath)
		if err != nil {
			logger.Printf("API component file not found: %s\n", filePath)
			http.NotFound(w, r)
			return
		}

		logger.Printf("Serving API component: %s\n", filePath)

		tmpl := template.Must(template.ParseFiles(filePath))
		tmpl.Execute(w, nil)
	})

	logger.Printf("Server started on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
