package helpers

import (
	"pc-monitoring/models"
	"pc-monitoring/monitors"
)

func GetStats() models.Response {
	cpu_list   := monitors.CPU()
	ram_usage  := monitors.MEM()
	disks_list := monitors.Disks()
	gpu_list   := monitors.GPU()

	return models.Response{
		CPU:  cpu_list,
		RAM:  ram_usage,
		Disk: disks_list,
		GPU:  gpu_list,
	}
}