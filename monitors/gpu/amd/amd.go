//go:build linux

package amd

// #cgo LDFLAGS: -lamd_smi
// #include <amd_smi/amd_smi/amdsmi.h>
import "C"

// Use this package to handles the dlopen bindings
// under the hood so it only looks for the
// shared library at runtime:
// https://github.com/hhk7734/amdsmi.go

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"pc-monitoring/models"
	"strconv"
	"strings"
	// "github.com/ROCm/rocm-systems/projects/amdsmi"
)

// TODO: Uptade to amdsmi package
// For now, it's too large do download

type Monitor struct{
    handles []C.amdsmi_processor_handle_t
    card    string
	count   int8
    ROCm    bool
}

func New() (*Monitor, error) {
	instance := &Monitor{
		count: 0,
        ROCm: false,
	}

    if instance.ROCm {
        err := instance.startROCm()

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

func (m *Monitor) Close() error {
    if m.ROCm {
        err := C.amdsmi_shut_down()
        
        if err != C.AMDSMI_STATUS_SUCCESS {
            return fmt.Errorf(
                "amdsmi_shut_down: %d",
                err,
            )
        }
    } else {
        println("Stop AMD Linux non-ROCm")
    }

	return nil
}

func (m *Monitor) Stats() ([]*models.GPUData, error) {
	var gpus []*models.GPUData
	var count uint8

	handles := make([]C.amdsmi_processor_handle, count)

	for _, handle := range handles {
		fmt.Println(handle)
		// var info C.amdsmi_asic_info_t
		var activity C.amdsmi_engine_usage_t

		// C.amdsmi_get_gpu_asic_info(handle, &info)

		gpus = append(gpus, &models.GPUData{
			Load: float64(C.amdsmi_get_gpu_activity(handle, &activity)),
		})

		// C.amdsmi_get_gpu_activity()
		// C.amdsmi_get_gpu_vram_usage()
		// C.amdsmi_get_temp_metric()
		// C.amdsmi_get_power_info()
	}

    return gpus, nil
}

func (m *Monitor) startROCm() error {
	err := C.amdsmi_init(0)

    if err != C.AMDSMI_STATUS_SUCCESS {
		err := fmt.Sprintf("error to Init SMI Header -> Error Code: %d", err)

        return err
    }

    return err
}

func (m *Monitor) startFileSys() error {
    entries, err := filepath.Glob(
		"/sys/class/drm/card*/device/driver",
	)

	if err != nil {
		return err
	}

	var result []*Monitor

	for _, driverPath := range entries {
		target, err := os.Readlink(driverPath)

		if err != nil {
			continue
		}

		if !strings.Contains(target, "amdgpu") {
			continue
		}

		card := filepath.Dir(
			filepath.Dir(driverPath),
		)

		result = append(result, &Monitor{
			card: card,
		})
	}

	return nil
}

func (m *Monitor) Refresh() (models.GPUData, error) {
	g := models.GPUData{
		Vendor: models.AMD,
	}

	name := m.readText(
		filepath.Join(m.card, "device", "product_name"),
	)

	if name == "" {
		name = m.readText(
			filepath.Join(m.card, "device", "product"),
		)
	}

	g.Name = name

	load, ok := m.readFloat(
		filepath.Join(
			m.card,
			"device",
			"gpu_busy_percent",
		),
	)
    
    if ok {
		g.Load = load
	}

	if temp, ok := m.readTemperature(m.card); ok {
		g.Temperature = temp
	}

	return g, nil
}

func (m *Monitor) readText(path string) string {
	data, err := os.ReadFile(path)

	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(data))
}

func (m *Monitor) readFloat(path string) (float64, bool) {
	data, err := os.ReadFile(path)
	if err != nil {
		return 0, false
	}

	v, err := strconv.ParseFloat(
		strings.TrimSpace(string(data)),
		64,
	)

	return v, err == nil
}

func (m *Monitor) readTemperature(card string) (float64, bool) {
	matches, _ := filepath.Glob(
		filepath.Join(
			card,
			"device",
			"hwmon",
			"hwmon*",
			"temp*_input",
		),
	)

	for _, path := range matches {
		v, ok := m.readFloat(path)
		if ok {
			return v / 1000.0, true
		}
	}

	return 0, false
}
