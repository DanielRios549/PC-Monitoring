//go:build windows

package nvidia

import (
    "fmt"
	// "unsafe"

    "golang.org/x/sys/windows"
	"pc-monitoring/models"
)

var nvmlPaths = []string{
    `C:\Windows\System32\nvml.dll`,
    `C:\Program Files\NVIDIA Corporation\NVSMI\nvml.dll`,
}

type NVML struct {
	dll *windows.DLL

	init                  *windows.Proc
	shutdown              *windows.Proc
	deviceGetCount        *windows.Proc
	deviceGetHandleByIdx  *windows.Proc
	deviceGetName         *windows.Proc
	deviceGetMemoryInfo   *windows.Proc
	deviceGetUtilization  *windows.Proc
	deviceGetTemperature  *windows.Proc
	deviceGetPowerUsage   *windows.Proc
	deviceGetPowerLimit   *windows.Proc
}

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
    stats := make([]*models.GPUData, 0)

    var dll *windows.DLL
	var err error

	for _, path := range nvmlPaths {
		dll, err = windows.LoadDLL(path)

		if err == nil {
			break
		}
	}

	if dll == nil {
		return nil, fmt.Errorf("nvml.dll not found: %w", err)
	}

	nvidia := &NVML{
		dll: dll,
	}

	nvidia.init, err = dll.FindProc("nvmlInit_v2")
	if err != nil {
		return nil, err
	}

	// nvidia.shutdown, err = dll.FindProc("nvmlShutdown")
	// if err != nil {
	// 	return nil, err
	// }

	// nvidia.deviceGetCount, err = dll.FindProc("nvmlDeviceGetCount_v2")
	// if err != nil {
	// 	return nil, err
	// }

	// nvidia.deviceGetHandleByIdx, err = dll.FindProc("nvmlDeviceGetHandleByIndex_v2")
	// if err != nil {
	// 	return nil, err
	// }

	nvidia.deviceGetName, err = dll.FindProc("nvmlDeviceGetName")
	if err != nil {
		return nil, err
	}

    stats = append(stats, &models.GPUData{
        Name: nvidia.deviceGetName.Name,
    })

    return stats, nil
}
