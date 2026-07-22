package monitors

import (
	"github.com/shirou/gopsutil/v3/cpu"

	"pc-monitoring/models"
)

// CPU usage per core
func CPU() []*models.CPUData {
	cpu_raw, _ := cpu.Percent(0, true)
	var cpu_list []*models.CPUData

	for i, v := range cpu_raw {
		cpu_list = append(cpu_list, &models.CPUData{
			Core: i,
			Usage: float64(v),
		})
	}

	return cpu_list
}
