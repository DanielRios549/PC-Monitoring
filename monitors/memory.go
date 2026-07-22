package monitors

import (
	"github.com/shirou/gopsutil/v3/mem"

	"pc-monitoring/models"
)

// TODO: Add Swap

// RAM usage
func MEM() *models.RAMData {
	vm, _ := mem.VirtualMemory()

	return &models.RAMData{
		Total_gb: float64(vm.Total) / (1024 * 1024 * 1024),
		Used_gb:  float64(vm.Used) / (1024 * 1024 * 1024),
		Percent:  float64(vm.UsedPercent),
	}
}
