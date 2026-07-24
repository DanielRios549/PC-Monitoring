package monitors

import (
	"github.com/shirou/gopsutil/v3/disk"

	"pc-monitoring/helpers"
	"pc-monitoring/models"
)

// TODO: Implement IO usage
// Disks Space
func Disks() []*models.DiskData {
	var disks_list []*models.DiskData
	partitions, _ := disk.Partitions(false)

	for _, p := range partitions {
		disk, _ := disk.Usage(p.Mountpoint)

		disks_list = append(disks_list, &models.DiskData{
			Device:  p.Device,
			Mount:   p.Mountpoint,
			Total:   helpers.RoundTo(float64(disk.Total) / (1024 * 1024 * 1024), 2),
			Used:    helpers.RoundTo(float64(disk.Used) / (1024 * 1024 * 1024), 2),
			Percent: helpers.RoundTo(float64(disk.UsedPercent), 2),
		})
	}


	return disks_list
}