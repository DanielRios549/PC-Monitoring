//go:build linux

package amd

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"pc-monitoring/helpers"
	"pc-monitoring/models"
	"strconv"
	"strings"
	"time"
)

// TODO: Uptade to amdsmi package
// For now, it's too large do download

type Monitor struct{
    devices []string
}

func New() (*Monitor, error) {
	instance := &Monitor{
		devices: helpers.FindPCIByVendor("0x1002"),
	}

    _, err := instance.Refresh()

    if err == nil {
        return instance, nil
    }

    return nil, errors.New("cannot Start AMD Linux. Neither FileSys or ROCm is working")
}

func (m *Monitor) CountDevices() int {
	return len(m.devices)
}

func (m *Monitor) Close() error {
    return nil
}

func (m *Monitor) Refresh() ([]*models.GPUData, error) {
	result := make([]*models.GPUData, 0, len(m.devices))

	for _, pci := range m.devices {
		g := readAMD(pci)

		if g.Name == "" {
			continue
		}

		result = append(result, g)
	}

	return result, nil
}

func readAMD(pci string) *models.GPUData {
	device := pci

	g := &models.GPUData{
		ID:        helpers.PciAddress(pci),
		Vendor:    models.AMD,
		Driver:    helpers.DriverForPCI(pci),
		Timestamp: time.Now(),
		// PCI:       helpers.PciAddress(pci),
		// Engines:   make(map[string]float64),
	}

	g.Name = amdName(device)

	/*
	 * amdgpu exposes gpu_busy_percent directly.
	 */
	g.Load = helpers.ReadFloat(
		filepath.Join(
			device,
			"gpu_busy_percent",
		),
	)

	g.Mem_used = uint64(helpers.ReadFloat(
		filepath.Join(
			device,
			"mem_busy_percent",
		),
	))

	g.Mem_total = helpers.ReadUint(
		filepath.Join(
			device,
			"mem_info_vram_total",
		),
	)

	g.Mem_used = helpers.ReadUint(
		filepath.Join(
			device,
			"mem_info_vram_used",
		),
	)

	/*
	 * Fallback to hwmon for temperature.
	 */
	g.Temperature = amdTemperature(device)

	/*
	 * Read clocks where exposed.
	 */
	g.CoreClock = amdClock(
		filepath.Join(
			device,
			"pp_dpm_sclk",
		),
	)

	g.MemoryClock = amdClock(
		filepath.Join(
			device,
			"pp_dpm_mclk",
		),
	)

	return g
}

func amdName(device string) string {
	for _, file := range []string{
		"product_name",
		"product_number",
	} {
		value := helpers.ReadString(
			filepath.Join(device, file),
		)

		if value != "" {
			return value
		}
	}

	/*
	 * sysfs doesn't always contain a friendly name.
	 */
	return strings.TrimSpace(
		helpers.ReadString(
			filepath.Join(device, "uevent"),
		),
	)
}

func amdTemperature(device string) float64 {
	matches, _ := filepath.Glob(
		filepath.Join(
			device,
			"hwmon",
			"hwmon*",
			"temp*_input",
		),
	)

	for _, path := range matches {
		value := helpers.ReadUint(path)

		if value > 0 {
			return float64(value) / 1000
		}
	}

	return 0
}

func amdClock(path string) uint64 {
	file, err := os.Open(path)
	if err != nil {
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if !strings.Contains(line, "*") {
			continue
		}

		fields := strings.Fields(line)

		for _, field := range fields {
			field = strings.TrimSuffix(
				field,
				"*",
			)

			if strings.HasSuffix(field, "Mhz") {
				value := strings.TrimSuffix(
					field,
					"Mhz",
				)

				n, err := strconv.ParseUint(
					value,
					10,
					64,
				)

				if err == nil {
					return n
				}
			}
		}
	}

	return 0
}
