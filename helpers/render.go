package helpers

import (
	"net/http"
	"html/template"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, file string, data interface{}) {
	tmplPath := filepath.Join("templates", file)
	tmpl, err := template.ParseFiles(tmplPath)

	if err != nil {
		http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)

	if err != nil {
		http.Error(w, "Template rendering error", http.StatusInternalServerError)
		return
	}
}
