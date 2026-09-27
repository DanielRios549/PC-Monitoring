package main

import (
	"pc-monitoring/modules"
	"pc-monitoring/functions"
)

func main() {
	functions.LoadEnv()

	server := modules.NewServer()
	tray := modules.NewTray(server)

	// Run WebServer in Parallel
	go server.Start()

	tray.Show()
}
