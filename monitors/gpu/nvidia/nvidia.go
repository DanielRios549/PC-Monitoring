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
            Mem_free:    mem.Total - mem.Used,
            // Temperature: float64(temp),
            // Power:       float64(power) / 1000,
        })
    }

    return stats, nil
}

func (p *Monitor) Refresh() ([]*models.GPUData, error) {
	count, ret := nvml.DeviceGetCount()
	if ret != nvml.SUCCESS {
		return nil, fmt.Errorf("DeviceGetCount: %s", nvml.ErrorString(ret))
	}

	result := make([]*models.GPUData, 0, count)

	for i := 0; i < count; i++ {
		dev, ret := nvml.DeviceGetHandleByIndex(i)
		if ret != nvml.SUCCESS {
			continue
		}

		name, _ := dev.GetName()
		uuid, _ := dev.GetUUID()
		driver, _ := nvml.SystemGetDriverVersion()

		g := &models.GPUData{
			ID:     uuid,
			Name:   name,
			Vendor: models.NVIDIA,
			Driver: driver,
		}

		if util, ret := dev.GetUtilizationRates(); ret == nvml.SUCCESS {
			g.Load = float64(util.Gpu)
		}

		if mem, ret := dev.GetMemoryInfo(); ret == nvml.SUCCESS {
			g.Mem_total = mem.Total
			g.Mem_used = mem.Used
			g.Mem_free = mem.Total - mem.Used

            if mem.Total > 0 {
                g.Mem_percent = float64(mem.Used) / float64(mem.Total) * 100.0
            }
		}

		if temp, ret := dev.GetTemperature(
			nvml.TEMPERATURE_GPU,
		); ret == nvml.SUCCESS {
			g.Temperature = float64(temp)
		}

		if power, ret := dev.GetPowerUsage(); ret == nvml.SUCCESS {
			g.PowerUsage = float64(power) / 1000.0
		}

		if limit, ret := dev.GetPowerManagementLimit(); ret == nvml.SUCCESS {
			g.PowerLimit = float64(limit) / 1000.0
		}

		if clock, ret := dev.GetClockInfo(
			nvml.CLOCK_GRAPHICS,
		); ret == nvml.SUCCESS {
			g.CoreClock =  uint64(clock)
		}

		if clock, ret := dev.GetClockInfo(
			nvml.CLOCK_MEM,
		); ret == nvml.SUCCESS {
			g.MemoryClock = uint64(clock)
		}

		if fan, ret := dev.GetFanSpeed(); ret == nvml.SUCCESS {
			g.FanSpeed = float64(fan)
		}

		result = append(result, g)
	}

	return result, nil
}
