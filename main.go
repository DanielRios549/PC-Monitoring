package main

import (
	"embed"
	"pc-monitoring/modules"
	"pc-monitoring/functions"
)

//go:embed templates/* templates/*/*
var templatesFS embed.FS

func main() {
	functions.LoadEnv()
	functions.LoadConfig("data/floor.json")

	server := modules.NewServer(templatesFS)
	tray := modules.NewTray(server)

	// Run WebServer in Parallel
	go server.Start()

	tray.Show()
}
