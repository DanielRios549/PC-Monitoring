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

	monitors, _, err := detectVendor()
	// println("1: ", len(monitors))
	// test, _ := monitors[0].Stats()
	// println("2: ", len(test))

	if err != nil {
        return gpu_list
    }

	for _, data := range monitors {
		stats, err := data.Refresh()

		if err != nil {
			continue
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
	}

	return gpu_list
}

func detectVendor() ([]models.GPUMonitor, int, error)  {
	var monitors []models.GPUMonitor
	var monitor  models.GPUMonitor
	var count    int
	var err      error

	// Nvidia
	monitor, err = nvidia.New()
	count += monitor.CountDevices()
	
	if err == nil && count >= 1 {
		monitors = append(monitors, monitor)
	}

	// AMD
	monitor, err = amd.New()
	count += monitor.CountDevices()
	
	if err == nil && count >= 1 {
		monitors = append(monitors, monitor)
	}

	// Intel
	monitor, err = intel.New()
	count += monitor.CountDevices()
	
	if err == nil && count >= 1 {
		monitors = append(monitors, monitor)
	}

	// Return the monitors list if exists
	if len(monitors) < 1 {
		return nil, 0, errors.New("no supported GPU found")
	}

	return monitors, count, nil

}
