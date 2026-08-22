package monitors

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/shirou/gopsutil/v4/disk"

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
			Model:   getDiskModel(partition.Device),
			Total:   helpers.RoundTo(float64(usage.Total) / (1024 * 1024 * 1024), 2),
			Used:    helpers.RoundTo(float64(usage.Used) / (1024 * 1024 * 1024), 2),
			Percent: helpers.RoundTo(float64(usage.UsedPercent), 2),
		})
	}


	return disks_list
}

func getDiskModel(dev string) string {
	partSuffix := regexp.MustCompile(`p?\d+$`)

	base := filepath.Base(dev)
	base = partSuffix.ReplaceAllString(base, "")

	// TODO: Add Windows support too
	modelPath := filepath.Join("/sys/block", base, "device/model")

	data, err := os.ReadFile(modelPath)

	if err != nil {
		println("<<>> Err: ", err.Error())
		return "Unknown"
	}

	return strings.TrimSpace(string(data))
}
