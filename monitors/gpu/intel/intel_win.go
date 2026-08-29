//go:build windows

package intel

import (
	"pc-monitoring/models"
)

type Monitor struct{
	devices []string
}

func New() (*Monitor, error) {
	instance := &Monitor{
		devices: make([]string, 0),
	}

    return instance, nil
}

func (m *Monitor) CountDevices() int {
	return len(m.devices)
}

func (m *Monitor) Close() error {
    return nil
}

func (m *Monitor) Refresh() ([]*models.GPUData, error) {
	var gpus []*models.GPUData

    return gpus, nil
}