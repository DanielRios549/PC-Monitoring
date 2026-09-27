package helpers

import (
	t "pc-monitoring/templates"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"os/exec"
	"path/filepath"
	"runtime"
)

var templates *template.Template

func InitTemplates() {
	var paths []string

	err := fs.WalkDir(t.TemplatesFS, ".", func(path string, d fs.DirEntry, err error) error {
        if err != nil {
			print(err.Error())
            return err
        }

        if !d.IsDir() && filepath.Ext(path) == ".html" {
            println("PATH: ", path)
            paths = append(paths, path)
        }

        return nil
    })

	if err != nil {
		fmt.Printf("Template read error: %v", err)
	}

	templates = template.Must(template.ParseFS(t.TemplatesFS, paths...))
}

func RenderTemplate(w http.ResponseWriter, file string, data any) {
	err := templates.ExecuteTemplate(w, file, data)

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
