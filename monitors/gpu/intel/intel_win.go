//go:build windows

package intel

import (
	"pc-monitoring/models"
)

type Monitor struct{
	count int8
}

func New() (*Monitor, error) {
	instance := &Monitor{
		count: 0,
	}

    return instance, nil
}

func (m *Monitor) CountDevices() int8 {
	return m.count
}

func (m *Monitor) Close() {}

func (m *Monitor) Stats() ([]*models.GPUData, error) {
	var gpus []*models.GPUData

    return gpus, nil
}