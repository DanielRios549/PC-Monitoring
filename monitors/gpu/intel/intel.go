package intel

// #cgo LDFLAGS: -lze_loader
// #include <level_zero/ze_api.h>
import "C"

import (
	"errors"
	"fmt"
	"pc-monitoring/models"
)

type Monitor struct{
	count int8
}

func New() (*Monitor, error) {
	instance := &Monitor{
		count: 0,
	}

	init := C.zeInit(0)

    if init != C.ZE_RESULT_SUCCESS {
		err := fmt.Sprintf("error to Init LevelZero Header -> Error Code: %x", init)
        return instance, errors.New(err)
    }

    return instance, nil
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