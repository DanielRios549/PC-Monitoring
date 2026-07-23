package main

import (
	"pc-monitoring/modules"
)

func main() {
	server := modules.NewServer()
	tray := modules.NewTray(server)

	// Run WebServer in Parallel
	go server.Start()

	tray.Show()
}
