package gpu

import (
	"errors"
	"pc-monitoring/models"
	"pc-monitoring/monitors/gpu/amd"
	"pc-monitoring/monitors/gpu/intel"
	"pc-monitoring/monitors/gpu/nvidia"
)

// GPU details
func GPU() []*models.GPUData {
	gpu_list := []*models.GPUData{}

	data, err := detectVendor()

	if err != nil {
        return gpu_list
    }

	stats, err := data.Stats()

	if err != nil {
        return gpu_list
    }

	for _, stat := range stats {
		gpu_list = append(gpu_list, &models.GPUData{
			Name:        stat.Name,
			Type:        stat.Type,
			Load:        stat.Load,
			Mem_total:   stat.Mem_total,
			Mem_used:    stat.Mem_used,
			Mem_percent: stat.Mem_percent,
		})
	}

	return gpu_list
}

func detectVendor() (models.GPUMonitor, error)  {
    if monitor, err := nvidia.New(); err == nil {
        return monitor, nil
    }

	if monitor, err := amd.New(); err == nil {
        return monitor, nil
    }

    if monitor, err := intel.New(); err == nil {
        return monitor, nil
    }

    return nil, errors.New("no supported GPU found")
}
