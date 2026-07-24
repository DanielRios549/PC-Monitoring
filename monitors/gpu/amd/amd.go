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
	"pc-monitoring/models"

	// "github.com/ROCm/rocm-systems/projects/amdsmi"

)

// TODO: Uptade to amdsmi package
// For now, it's too large do download

type Monitor struct{
	count int8
}

func New() (*Monitor, error) {
	instance := &Monitor{
		count: 0,
	}

	init := C.amdsmi_init(0)

    if init != C.AMDSMI_STATUS_SUCCESS {
		err := fmt.Sprintf("error to Init SMI Header -> Error Code: %d", init)
        return instance, errors.New(err)
    }

    return instance, nil
}

func (m *Monitor) CountDevices() int8 {
	return m.count
}

func (m *Monitor) Close() {
    C.amdsmi_shut_down()
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
