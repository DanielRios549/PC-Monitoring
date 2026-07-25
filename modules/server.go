package modules

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"pc-monitoring/helpers"
	"pc-monitoring/models"
	"pc-monitoring/monitors"
	"pc-monitoring/monitors/gpu"

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
		cpu := monitors.CPUInfo()

		header := models.PageData{
			Title:       "PC Monitoring",
			Description: "Real-time PC monitoring tool",
		}

		info := &models.Response{
			CPU: cpu,
		}

		pageInfo := map[string]any{
			"Header": header,
			"Info": info,
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		helpers.RenderTemplate(w, "index.html", pageInfo)
	})

	s.router.Post("/cpu", func(w http.ResponseWriter, r *http.Request) {
		monitorData := monitors.CPUData()
	
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		helpers.RenderTemplate(w, "cpu.html", monitorData)
	})

	s.router.Post("/memory", func(w http.ResponseWriter, r *http.Request) {
		monitorData := monitors.MEMData()
	
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		helpers.RenderTemplate(w, "memory.html", monitorData)
	})

	s.router.Post("/disk", func(w http.ResponseWriter, r *http.Request) {
		monitorData := monitors.Disks()
	
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		helpers.RenderTemplate(w, "disk.html", monitorData)
	})

	s.router.Post("/gpu", func(w http.ResponseWriter, r *http.Request) {
		monitorData := gpu.GPU()
	
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		helpers.RenderTemplate(w, "gpu.html", monitorData)
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
