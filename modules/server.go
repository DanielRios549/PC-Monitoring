package modules

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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

	instance := &Server{
		server: nil,
		router: router,
	}

	instance.Routes()

	return instance
}

func (s *Server) Routes() {
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

	if s.server != nil {
		fmt.Println("Server already running")
		return
	}

	s.server = &http.Server{
		Addr: ":9003",
		Handler: s.router,
	}

	err := s.server.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		fmt.Printf("Error to Start Web Server: %v\n", err)
		return
	}

	helpers.ServerStatus <- true
}

func (s *Server) Stop() {
	fmt.Println("Stopping Monitoring server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := s.server.Shutdown(ctx)
	
	if err != nil {
		// Close Server if Gracefully Shutdown fails
		err = s.server.Close()

		if err != nil {
			fmt.Printf("Error to Stop Web Server: %v\n", err)
			return
		}

		return
	}

	s.server = nil
	helpers.ServerStatus <- false
}
