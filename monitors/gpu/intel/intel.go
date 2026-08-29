//go:build linux

package intel

import (
	"errors"
	"os"
	"path/filepath"
	"pc-monitoring/helpers"
	"pc-monitoring/models"
	"strconv"
	"strings"
	"time"
)

type Monitor struct{
    devices     []string
}

func New() (*Monitor, error) {
    var err error

	instance := &Monitor{
		devices: helpers.FindPCIByVendor("0x8086"),
	}

    _, err = instance.Refresh()

    if err == nil {
        return instance, nil
    }

    return nil, errors.New("cannot Start AMD on Linux. Neither FileSys or ROCm is working")
}

func (m *Monitor) CountDevices() int {
	return len(m.devices)
}

func (m *Monitor) Close() error {
    return nil
}

func (m *Monitor) Refresh()  ([]*models.GPUData, error)  {
	result := make([]*models.GPUData, 0, len(m.devices))

	for _, pci := range m.devices {
		driver := helpers.DriverForPCI(pci)

		if driver != "i915" && driver != "xe" {
			continue
		}

		g := readIntel(pci)

		result = append(result, g)
	}

	return result, nil
}

func (m *Monitor) IntelUtilization(path string) (float64, bool) {
	/*
	 * Different Intel generations expose different
	 * sysfs/DRM files.
	 *
	 * Try common interfaces first.
	 */

	candidates := []string{
		filepath.Join(path, "gpu_busy_percent"),
		filepath.Join(path, "gt", "busy_percent"),
	}

	for _, candidate := range candidates {
		data, err := os.ReadFile(candidate)
		if err != nil {
			continue
		}

		v, err := strconv.ParseFloat(
			strings.TrimSpace(string(data)),
			64,
		)

		if err == nil {
			return v, true
		}
	}

	return 0, false
}

func (m *Monitor) ReadText(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(data))
}

func readIntel(pci string) *models.GPUData {
	g := &models.GPUData{
		ID:        helpers.PciAddress(pci),
		Vendor:    models.Intel,
		Driver:    helpers.DriverForPCI(pci),
		// PCI:       helpers.PciAddress(pci),
		Timestamp: time.Now(),
		// Engines:   make(map[string]float64),
	}

	g.Name = intelName(pci)

	/*
	 * Intel shared/system GPU memory.
	 */
	g.Mem_total = intelMemoryTotal(pci)

	/*
	 * Temperature through hwmon if available.
	 */
	g.Temperature = intelTemperature(pci)

	return g
}

func intelName(pci string) string {
	/*
	 * Prefer the PCI modalias/device information.
	 *
	 * A friendly marketing name is not guaranteed by sysfs.
	 */
	deviceID := helpers.ReadString(
		filepath.Join(pci, "device"),
	)

	if deviceID != "" {
		return "Intel GPU " + deviceID
	}

	return "Intel GPU"
}

func intelTemperature(device string) float64 {
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

func intelMemoryTotal(pci string) uint64 {
	/*
	 * Intel integrated GPUs generally use system memory.
	 *
	 * i915 exposes:
	 *   /sys/class/drm/cardX/device/mem_info_stolen_total
	 *
	 * xe can expose:
	 *   /sys/class/drm/cardX/device/tileX/mem_info_vram_total
	 *
	 * For integrated GPUs, use the system's total RAM as the
	 * maximum addressable GPU memory when no dedicated VRAM
	 * information is available.
	 */

	driver := helpers.DriverForPCI(pci)

	switch driver {
	case "xe":
		if value := helpers.ReadUint(
			filepath.Join(
				pci,
				"mem_info_vram_total",
			),
		); value > 0 {
			return value
		}

		// Try tile-based layout.
		matches, _ := filepath.Glob(
			filepath.Join(
				pci,
				"tile*",
				"mem_info_vram_total",
			),
		)

		var total uint64

		for _, path := range matches {
			total += helpers.ReadUint(path)
		}

		if total > 0 {
			return total
		}

	case "i915":
		/*
		 * i915 stolen memory is not the same thing as total
		 * GPU-accessible system memory, so don't report it
		 * as total VRAM.
		 */
	}

	/*
	 * Intel integrated GPU:
	 *
	 * GPU memory is shared with system RAM.
	 *
	 * /proc/meminfo gives us total system memory.
	 */
	return systemMemoryTotal()
}

func systemMemoryTotal() uint64 {
	data := helpers.ReadString("/proc/meminfo")

	for _, line := range strings.Split(data, "\n") {
		fields := strings.Fields(line)

		if len(fields) < 2 {
			continue
		}

		if fields[0] != "MemTotal:" {
			continue
		}

		value, err := strconv.ParseUint(
			fields[1],
			10,
			64,
		)

		if err != nil {
			return 0
		}

		// /proc/meminfo reports kB.
		return value * 1024
	}

	return 0
}
