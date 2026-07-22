package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"pc-monitoring/models"
	"pc-monitoring/helpers"
)

func main() {
	router := chi.NewRouter()

	fileServer := http.FileServer(http.Dir("./static"))
	router.Handle("/*", http.StripPrefix("/", fileServer))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		pageData := models.PageData{
			Title:   "PC Monitoring",
			Message: "Real-time PC monitoring tool",
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		helpers.RenderTemplate(w, "index.html", pageData)
	})

	router.Get("/stats", func(w http.ResponseWriter, r *http.Request) {
		monitorData := helpers.GetStats()
	
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(monitorData)

		if err != nil {
			http.Error(w, "Error to Show Stats Data", http.StatusInternalServerError)
		}
	})

	fmt.Println("Server starting on :9003...")
	err := http.ListenAndServe(":9003", router)

	if err != nil {
		fmt.Printf("Error to Start the Application: %v", err)
	}
}
