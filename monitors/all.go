package monitors

import (
	"pc-monitoring/models"
)

func GetStats() models.Response {
	cpu_list   := CPU()
	ram_list   := MEM()
	disks_list := Disks()
	gpu_list   := GPU()

	return models.Response{
		CPU:  cpu_list,
		RAM:  ram_list,
		Disk: disks_list,
		GPU:  gpu_list,
	}
}