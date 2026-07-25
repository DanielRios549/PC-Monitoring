package helpers

import (
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
	"path/filepath"
	"runtime"
)

func RenderTemplate(w http.ResponseWriter, file string, data any) {
	tmplPath := filepath.Join("templates", file)
	tmpl, err := template.ParseFiles(tmplPath)
	
	if err != nil {
		message := fmt.Sprintf("Template parsing error: %v", err.Error())
		http.Error(w, message, http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)

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
