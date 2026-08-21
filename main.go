package main

import (
	"embed"
	"pc-monitoring/modules"
)

//go:embed templates/* templates/*/*
var templatesFS embed.FS

func main() {
	server := modules.NewServer(templatesFS)
	tray := modules.NewTray(server)

	// Run WebServer in Parallel
	go server.Start()

	tray.Show()
}
