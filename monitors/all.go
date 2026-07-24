package monitors

import (
	"pc-monitoring/models"
	"pc-monitoring/monitors/gpu"
)

func GetStats() models.Response {
	cpu_list   := CPU()
	ram_list   := MEM()
	disks_list := Disks()
	gpu_list   := gpu.GPU()

	return models.Response{
		CPU:  cpu_list,
		RAM:  ram_list,
		Disk: disks_list,
		GPU:  gpu_list,
	}
}