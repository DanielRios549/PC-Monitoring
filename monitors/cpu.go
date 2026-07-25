package monitors

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/cpu"

	"pc-monitoring/helpers"
	"pc-monitoring/models"
)

// CPU Information
func CPUInfo() *models.CPUInfo {
	getAll, err := cpu.Info()

	if err != nil || len(getAll) < 1 {
		fmt.Printf("Error to get CPU Information: %v", err)

		return &models.CPUInfo{}
	}

	get := getAll[0]

	info := &models.CPUInfo{
		Name: get.ModelName,
		CoreCount: len(getAll),
	}

	return info
}

// Real-time CPU usage per thread
func CPUData() []*models.CPUData {
	cpu_raw, _ := cpu.Percent(0, true)
	var cpu_list []*models.CPUData

	for core, usage := range cpu_raw {
		usage := helpers.RoundTo(usage, 2)
	
		cpu_list = append(cpu_list, &models.CPUData{
			Core:  core + 1,
			Usage: usage,
		})
	}

	return cpu_list
}
