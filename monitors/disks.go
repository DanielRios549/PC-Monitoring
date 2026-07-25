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

	for _, partition := range partitions {
		usage, _ := disk.Usage(partition.Mountpoint)
		// speeds, _ := disk.IOCounters(partition.Device)

		// for _, oip := range speeds {
		// 	fmt.Println("Device ", partition.Device, " ", oip.ReadBytes)
		// }

		disks_list = append(disks_list, &models.DiskData{
			Device:  partition.Device,
			Mount:   partition.Mountpoint,
			Total:   helpers.RoundTo(float64(usage.Total) / (1024 * 1024 * 1024), 2),
			Used:    helpers.RoundTo(float64(usage.Used) / (1024 * 1024 * 1024), 2),
			Percent: helpers.RoundTo(float64(usage.UsedPercent), 2),
		})
	}


	return disks_list
}