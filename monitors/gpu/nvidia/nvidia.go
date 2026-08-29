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
	devices []string
}

func New() (*Monitor, error) {
	instance := &Monitor{
		devices: make([]string, 0),
	}

	init := nvml.Init()

    if init != nvml.SUCCESS {
		err := fmt.Sprintf("error to Init NVML -> Error Code: %v", init)
        return instance, errors.New(err)
    }

    return instance, nil
}

func (m *Monitor) CountDevices() int {
	return len(m.devices)
}

func (m *Monitor) Close() error {
    err := nvml.Shutdown()

	if err.Error() != "" {
		fmt.Println("Error to close NVML")
	}

    return err
}

func (m *Monitor) Refresh() ([]*models.GPUData, error) {
    count, ret := nvml.DeviceGetCount()

    if ret != nvml.SUCCESS {
        return nil, ret
    }

    stats := make([]*models.GPUData, 0, count)

    for index := range count {
        dev, _ := nvml.DeviceGetHandleByIndex(index)

        name, _ := dev.GetName()
		uuid, _ := dev.GetUUID()
        driver, _ := nvml.SystemGetDriverVersion()

        gpu := &models.GPUData{
			ID:     uuid,
			Name:   name,
			Vendor: models.NVIDIA,
			Driver: driver,
		}

		if util, ret := dev.GetUtilizationRates(); ret == nvml.SUCCESS {
			gpu.Load = helpers.RoundTo(float64(util.Gpu), 2)
		}

		if mem, ret := dev.GetMemoryInfo(); ret == nvml.SUCCESS {
			gpu.Mem_total = mem.Total
			gpu.Mem_used = mem.Used
			gpu.Mem_free = mem.Total - mem.Used

            if mem.Total > 0 {
                gpu.Mem_percent = float64(mem.Used) / float64(mem.Total) * 100
            }
		}

		if temp, ret := dev.GetTemperature(nvml.TEMPERATURE_GPU); ret == nvml.SUCCESS {
			gpu.Temperature = float64(temp)
		}

		if power, ret := dev.GetPowerUsage(); ret == nvml.SUCCESS {
			gpu.PowerUsage = float64(power) / 1000
		}

		if limit, ret := dev.GetPowerManagementLimit(); ret == nvml.SUCCESS {
			gpu.PowerLimit = float64(limit) / 1000
		}

		if clock, ret := dev.GetClockInfo(nvml.CLOCK_GRAPHICS); ret == nvml.SUCCESS {
			gpu.CoreClock =  uint64(clock)
		}

		if clock, ret := dev.GetClockInfo(nvml.CLOCK_MEM); ret == nvml.SUCCESS {
			gpu.MemoryClock = uint64(clock)
		}

		if fan, ret := dev.GetFanSpeed(); ret == nvml.SUCCESS {
			gpu.FanSpeed = float64(fan)
		}

		stats = append(stats, gpu)
    }

    return stats, nil
}
