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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		urlPath := r.URL.Path
		fmt.Printf("Request: %s\n", urlPath)
		if urlPath == "/" {
			urlPath = "/index"
		}
		filePath := filepath.Join(pagesDir, urlPath+".html")
		_, err := os.Stat(filePath)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		tmpl := template.Must(template.ParseFiles(filePath))
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/api/Box", func(w http.ResponseWriter, r *http.Request) {
		componetName := r.URL.Path[len("/api/"):]
		tmpl := template.Must(template.ParseFiles("componets/" + componetName + ".html"))
		tmpl.Execute(w, nil)
	})
	port := 3500
	fmt.Printf("Server started on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
