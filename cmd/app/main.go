package main

import (
	"pc-monitoring/modules"
	"pc-monitoring/functions"
)

func main() {
	functions.LoadEnv()
    functions.CheckCompanies()

	server := modules.NewServer()
    agent  := modules.NewAgent()
	tray   := modules.NewTray(server)

	// Run WebServer and Agent in Parallel
	go server.Start()
    go agent.Start()

	tray.Show()
}
