package monitors

import (
	"github.com/shirou/gopsutil/v3/disk"

	"pc-monitoring/models"
)

func Disks() []*models.DiskData {
	var disks_list []*models.DiskData
	partitions, _ := disk.Partitions(false)

	for _, p := range partitions {
		disk, _ := disk.Usage(p.Mountpoint)

		disks_list = append(disks_list, &models.DiskData{
			Device:  p.Device,
			Mount:   p.Mountpoint,
			Total:   float64(disk.Total) / (1024 * 1024 * 1024),
			Used:    float64(disk.Used) / (1024 * 1024 * 1024),
			Percent: float64(disk.UsedPercent),
		})
	}


	return disks_list
}