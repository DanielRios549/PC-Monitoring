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
		usage := helpers.RoundTo(usage, 2)
		color := "#334155"

		if usage > 90 {
			color = "#451a1a"
		}
	
		cpu_list = append(cpu_list, &models.CPUData{
			Core:  core + 1,
			Usage: usage,
			Color: color,
		})
	}

	return cpu_list
}
