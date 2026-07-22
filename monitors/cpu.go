package monitors

import (
	"github.com/shirou/gopsutil/v3/cpu"

	"pc-monitoring/models"
	"pc-monitoring/helpers"
)

// CPU usage per core
func CPU() []*models.CPUData {
	cpu_raw, _ := cpu.Percent(0, true)
	var cpu_list []*models.CPUData

	for core, usage := range cpu_raw {
		cpu_list = append(cpu_list, &models.CPUData{
			Core:  core,
			Usage: helpers.RoundTo(usage, 2),
		})
	}

	return cpu_list
}
