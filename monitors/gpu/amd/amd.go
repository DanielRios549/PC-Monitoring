package amd

// #cgo LDFLAGS: -lrocm_smi64
// #include <rocm_smi/rocm_smi.h>
import "C"

import (
	"errors"
	"pc-monitoring/models"
)

type Monitor struct{}

func New() (*Monitor, error) {
    if C.rsmi_init(0) != C.RSMI_STATUS_SUCCESS {
        return nil, errors.New("error to Init ROCm Header")
    }

    return &Monitor{}, nil
}

func (m *Monitor) Close() {
    C.rsmi_shut_down()
}

func (m *Monitor) Stats() ([]*models.GPUData, error) {
    // enumerate processors
    // call:
    //
    // amdsmi_get_gpu_activity()
    // amdsmi_get_gpu_vram_usage()
    // amdsmi_get_temp_metric()
    // amdsmi_get_power_info()
    //
    return nil, nil
}