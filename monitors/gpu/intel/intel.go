package intel

// #cgo LDFLAGS: -lze_loader
// #include <level_zero/ze_api.h>
import "C"

import (
	"errors"
	"pc-monitoring/models"
)

type Monitor struct{}

func New() (*Monitor, error) {
    if C.zeInit(0) != C.ZE_RESULT_SUCCESS {
        return nil, errors.New("error to Init LeverZero Header")
    }

    return &Monitor{}, nil
}

func (m *Monitor) Close() {}

func (m *Monitor) Stats() ([]*models.GPUData, error) {
    // enumerate drivers
    // enumerate devices
    // query metric groups
    // query memory
    // query power
    return nil, nil
}