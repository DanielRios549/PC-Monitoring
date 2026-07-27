package helpers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

var templates = make(map[string]*template.Template)

func InitTemplates() {
	entries, err := os.ReadDir("templates")

	if err != nil {
		fmt.Printf("Template parsing error: %v", err.Error())
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		file := entry.Name()

		path := filepath.Join("templates", file)
		templates[file] = template.Must(template.ParseFiles(path))
	}
}

func RenderTemplate(w http.ResponseWriter, file string, data any) {
	tmpl := templates[file]
	err := tmpl.Execute(w, data)

	if err != nil {
		message := fmt.Sprintf("Template rendering error: %v", err.Error())
		http.Error(w, message, http.StatusInternalServerError)
		return
	}
}

func OpenBrowser(url string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
		case "windows":
			cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
		case "darwin":
			cmd = exec.Command("open", url)
		default:
			cmd = exec.Command("xdg-open", url)
	}

	return cmd.Start()
}
