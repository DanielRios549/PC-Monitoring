//go:build linux

package nvidia

import (
	"errors"
	"fmt"
	"pc-monitoring/helpers"
	"pc-monitoring/models"

	"github.com/NVIDIA/go-nvml/pkg/nvml"
)

type Monitor struct{
	count int8
}

func New() (*Monitor, error) {
	instance := &Monitor{
		count: 0,
	}

	init := nvml.Init()

    if init != nvml.SUCCESS {
		err := fmt.Sprintf("error to Init NVML -> Error Code: %v", init)
        return instance, errors.New(err)
    }

    return instance, nil
}

func (m *Monitor) CountDevices() int8 {
	return m.count
}

func (m *Monitor) Close() {
    err := nvml.Shutdown()

	if err.Error() != "" {
		fmt.Println("Error to close NVML")
	}
}

func (m *Monitor) Stats() ([]*models.GPUData, error) {
    count, ret := nvml.DeviceGetCount()

    if ret != nvml.SUCCESS {
        return nil, ret
    }

    stats := make([]*models.GPUData, 0, count)

    for i := 0; i < count; i++ {
        dev, _ := nvml.DeviceGetHandleByIndex(i)

        name, _ := dev.GetName()
        util, _ := dev.GetUtilizationRates()
        mem, _ := dev.GetMemoryInfo()
        // temp, _ := dev.GetTemperature(nvml.TEMPERATURE_GPU)
        // power, _ := dev.GetPowerUsage()

        stats = append(stats, &models.GPUData{
            // Vendor:      "NVIDIA",
            Name:        name,
            Load:        helpers.RoundTo(float64(util.Gpu), 2),
            Mem_used:    mem.Used,
            Mem_total:   mem.Total,
            // Temperature: float64(temp),
            // Power:       float64(power) / 1000,
        })
    }

    return stats, nil
}
