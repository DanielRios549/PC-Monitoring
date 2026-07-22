package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"pc-monitoring/models"
	"pc-monitoring/helpers"
	"pc-monitoring/monitors"
)

func main() {
	router := chi.NewRouter()
	server := &http.Server{
		Addr: ":9003",
		Handler: router,
	}

	fileServer := http.FileServer(http.Dir("./static"))
	router.Handle("/*", fileServer)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		pageData := models.PageData{
			Title:   "PC Monitoring",
			Message: "Real-time PC monitoring tool",
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		helpers.RenderTemplate(w, "index.html", pageData)
	})

	router.Get("/stats", func(w http.ResponseWriter, r *http.Request) {
		monitorData := monitors.GetStats()
	
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(monitorData)

		if err != nil {
			http.Error(w, "Error to Show Stats Data", http.StatusInternalServerError)
			return
		}
	})

	fmt.Println("Monitoring server starting on :9003...")
	err := server.ListenAndServe()

	if err != nil {
		fmt.Printf("Error to Start the Application: %v", err)
		return
	}
}
