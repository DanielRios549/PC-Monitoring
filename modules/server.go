package modules

import (
	"encoding/json"
	"fmt"
	"net/http"

	"pc-monitoring/helpers"
	"pc-monitoring/models"
	"pc-monitoring/monitors"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	server  *http.Server
	router  *chi.Mux
}

func NewServer() *Server {
	router := chi.NewRouter()

	return &Server{
		server: &http.Server{
			Addr: ":9003",
			Handler: router,
		},
		router: router,
	}
}

func (s *Server) Router() {
	fileServer := http.FileServer(http.Dir("./static"))
	s.router.Handle("/*", fileServer)

	s.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		pageData := models.PageData{
			Title:   "PC Monitoring",
			Message: "Real-time PC monitoring tool",
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		helpers.RenderTemplate(w, "index.html", pageData)
	})

	s.router.Get("/stats", func(w http.ResponseWriter, r *http.Request) {
		monitorData := monitors.GetStats()
	
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(monitorData)

		if err != nil {
			http.Error(w, "Error to Show Stats Data", http.StatusInternalServerError)
			return
		}
	})
}

func (s *Server) Start() {
	fmt.Println("Starting Monitoring server at :9003...")
	s.Router()

	// done := helpers.InitWatch()
	err := s.server.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		fmt.Printf("Error to Start Web Server: %v\n", err)
		return
	}

	helpers.ServerStatus <- true
	// done <- true
}

func (s *Server) Stop() {
	fmt.Println("Stopping Monitoring server...")
	err := s.server.Close()
	
	if err != nil {
		fmt.Printf("Error to Stop Web Server: %v\n", err)
		return
	}

	helpers.ServerStatus <- false
}
