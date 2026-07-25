package monitors

import (
	"github.com/shirou/gopsutil/v3/mem"

	"pc-monitoring/models"
	"pc-monitoring/helpers"
)

// RAM and Swap usage
func MEMData() []*models.RAMData {
	var ram []*models.RAMData

	system, _ := mem.VirtualMemory()

	ram = append(ram, &models.RAMData{
		Total_gb: helpers.RoundTo(float64(system.Total) / (1024 * 1024 * 1024), 2),
		Used_gb:  helpers.RoundTo(float64(system.Used) / (1024 * 1024 * 1024), 2),
		Percent:  helpers.RoundTo(float64(system.UsedPercent), 2),
	})

	swap, err := mem.SwapMemory()

	if err == nil {
		ram = append(ram, &models.RAMData{
			Total_gb: helpers.RoundTo(float64(swap.Total) / (1024 * 1024 * 1024), 2),
			Used_gb:  helpers.RoundTo(float64(swap.Used) / (1024 * 1024 * 1024), 2),
			Percent:  helpers.RoundTo(float64(swap.UsedPercent), 2),
		})
	}

	return ram
}
