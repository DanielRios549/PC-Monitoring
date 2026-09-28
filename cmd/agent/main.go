package main

import (
	"pc-monitoring/modules"
	"pc-monitoring/functions"
)

func main() {
	functions.LoadEnv()

	agent  := modules.NewAgent()
	agent.Start()
}
