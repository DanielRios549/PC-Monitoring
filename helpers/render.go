package helpers

import (
	"net/http"
	"html/template"
	"path/filepath"
	"pc-monitoring/models"
)

func RenderTemplate(w http.ResponseWriter, file string, data models.PageData) {
	tmplPath := filepath.Join("templates", file)
	tmpl, err := template.ParseFiles(tmplPath)

	if err != nil {
		http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, file)

	if err != nil {
		http.Error(w, "Template rendering error", http.StatusInternalServerError)
		return
	}
}
