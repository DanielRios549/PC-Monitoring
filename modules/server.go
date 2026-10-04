package modules

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"pc-monitoring/functions"
	"pc-monitoring/helpers"
	"pc-monitoring/monitors"
	"pc-monitoring/monitors/gpu"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	server      *http.Server
	router      *chi.Mux
}

func NewServer() *Server {
	router := chi.NewRouter()

	instance := &Server{
		server: nil,
		router: router,
	}

	helpers.InitTemplates()
	instance.Routes()

	return instance
}

func (s *Server) Routes() {
	fileServer := http.FileServer(http.Dir("./static"))
	s.router.Handle("/*", fileServer)

    s.Pages()
    s.Monitors()
}

func (s *Server) Pages() {
    s.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
        pageData := functions.PageData("PC Monitoring", "Real-time PC monitoring tool")

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		helpers.RenderTemplate(w, "pages/index", pageData)
	})

	s.router.Get("/devices", func(w http.ResponseWriter, r *http.Request) {
        pageData := functions.PageData("Devices - Work In Progress", "Show available Devices")
    
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		helpers.RenderTemplate(w, "pages/devices", pageData)
	})

    s.router.Get("/printers", func(w http.ResponseWriter, r *http.Request) {
        pageData := functions.PageData("Printers", "Show available Printers")
        printers := functions.LoadPrinterConfig("data/floor.json")

        data := map[string]any{
            "Printers":  printers,
            "Data":      pageData,
        }

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		helpers.RenderTemplate(w, "pages/printers", data)
	})

    s.router.Get("/access-points", func(w http.ResponseWriter, r *http.Request) {
        pageData := functions.PageData("Access Points", "Show available APs")
        ap := functions.LoadAPConfig("data/floor.json")

        data := map[string]any{
            "APs":    ap,
            "Data":  pageData,
        }

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		helpers.RenderTemplate(w, "pages/access-points", data)
	})

	s.router.Get("/settings", func(w http.ResponseWriter, r *http.Request) {
        pageData := functions.PageData("Settings - Work In Progress", "System Settings")

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		helpers.RenderTemplate(w, "pages/settings", pageData)
	})
}

func (s *Server) Monitors() {
    s.router.Post("/cpu", func(w http.ResponseWriter, r *http.Request) {
		monitorData := monitors.CPUData()
	
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		helpers.RenderTemplate(w, "monitors/cpu", monitorData)
	})

	s.router.Post("/memory", func(w http.ResponseWriter, r *http.Request) {
		monitorData := monitors.MEMData()
	
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		helpers.RenderTemplate(w, "monitors/memory", monitorData)
	})

	s.router.Post("/disk", func(w http.ResponseWriter, r *http.Request) {
		monitorData := monitors.Disks()
	
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		helpers.RenderTemplate(w, "monitors/disk", monitorData)
	})

	s.router.Post("/gpu", func(w http.ResponseWriter, r *http.Request) {
		monitorData := gpu.GPU()
	
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		helpers.RenderTemplate(w, "monitors/gpu", monitorData)
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
