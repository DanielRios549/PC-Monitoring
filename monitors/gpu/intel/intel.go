//go:build linux

package intel

// #cgo LDFLAGS: -lze_loader
// #include <level_zero/ze_api.h>
import "C"

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"pc-monitoring/models"
	"strconv"
	"strings"
)

type Monitor struct{
    devices     []string
	count       int8
    level_zero  bool
}

func New() (*Monitor, error) {
	instance := &Monitor{
		count: 0,
        level_zero: false,
	}

    if instance.level_zero {
        err := instance.startLevelZero()

        if err == nil {
            return instance, nil
        }
    } else {
        err := instance.startFileSys()

        if err == nil {
            return instance, nil
        }
    }

    return nil, errors.New("Cannot Start AMD Linux. Neither FileSys or ROCm is working")
}

func (m *Monitor) CountDevices() int8 {
	return m.count
}

func (m *Monitor) Close() {}

func (m *Monitor) Stats() ([]*models.GPUData, error) {
	var gpus []*models.GPUData
	var driverCount C.uint32_t

	drivers := make([]C.ze_driver_handle_t, driverCount)
	C.zeDriverGet(&driverCount, &drivers[0])

	for _, driver := range drivers {
		var deviceCount C.uint32_t

		devices := make([]C.ze_device_handle_t, deviceCount)
		C.zeDeviceGet(driver, &deviceCount, &devices[0])

		for _, device := range devices {
			var props C.ze_device_properties_t
			props.stype = C.ZE_STRUCTURE_TYPE_DEVICE_PROPERTIES

			C.zeDeviceGetProperties(device, &props)
			

			fmt.Println(C.GoString(&props.name[0]))

			gpus = append(gpus, &models.GPUData{
				Name: C.GoString(&props.name[0]),
				// Load: float64(C.amdsmi_get_gpu_activity(handle, &activity)),
			})
		}
	}

    return gpus, nil
}

func (m *Monitor) startLevelZero() error {
    err := C.zeInit(0)

    if err != C.ZE_RESULT_SUCCESS {
		err := fmt.Sprintf("error to Init LevelZero Header -> Error Code: %x", err)
        return err
    }

    return  err
}

func (m *Monitor) startFileSys() error {
	paths, err := filepath.Glob(
		"/sys/class/drm/card*/device",
	)
	if err != nil {
		return err
	}

	var devices []string

	for _, path := range paths {
		vendor := m.readText(
			filepath.Join(path, "vendor"),
		)

		if vendor == "0x8086" {
			devices = append(devices, path)
		}
	}

    m.devices = devices

	return nil
}

func (m *Monitor) intelUtilization(path string) (float64, bool) {
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

func (m *Monitor) readText(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(data))
}
