//go:build windows

package amd

import (
	"fmt"
	"pc-monitoring/helpers"
	"pc-monitoring/models"
	"strings"

	"github.com/yusufpapurcu/wmi"
)

type Win32_Adapter struct {
	Name        string
	AdapterRAM  uint32 // TODO: Get true VRAM size
	PNPDeviceID string
}

type Win32_Performance struct {
	Name                  string
	UtilizationPercentage uint64
}

type Win32_PerfGPUMemory struct {
	Name         string
	CurrentUsage uint64
}

type GPUReport struct {
	Model       string
	VRAM        uint32
	LUID        string
	Utilization uint64
}

type Monitor struct{
	devices     []string
	controllers []Win32_Adapter
	engines     []Win32_Performance
	memories    []Win32_PerfGPUMemory
}

func New() (*Monitor, error) {
	instance := &Monitor{
		devices: make([]string, 0),
	}

	// Query GPU Information
	instance.getAdapter()
	instance.getEngines()
	instance.getMemories()

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

	type intermediate struct {
		Data       *models.GPUData
		EngineLoad uint64
	}

	gpuMap := make(map[string]*intermediate)
	foundAMD := false

	for _, ctrl := range m.controllers {
		nameLower := strings.ToLower(ctrl.Name)

		isAMD := strings.Contains(nameLower, "amd")
		isRadeon := strings.Contains(nameLower, "radeon")
	
		if isAMD || isRadeon {
			foundAMD = true

			luid := helpers.ExtractLUID(ctrl.PNPDeviceID)
			
			if luid == "" {
				continue
			}

			// vramBytes := helpers.GetTrueVRAM(ctrl.PNPDeviceID)
			// vramMB := vramBytes / (1024 * 1024)
			vramGB := ctrl.AdapterRAM / (1024 * 1024)
			gpuType := "Integrated"

			// TODO: Heuristic: GPUs with > 2GB dedicated VRAM are Discrete
			if vramGB > 2048 {
				gpuType = "Discrete"
			}

			gpuMap[luid] = &intermediate{
				Data: &models.GPUData{
					Name:      ctrl.Name,
					Type:      gpuType,
					Mem_total: uint64(vramGB),
				},
			}

			// gpus = append(gpus, &models.GPUData{
			// 	Name: gpu.Name,
			// 	Mem_total: uint64(helpers.RoundTo(float64(ramGB), 2)),
			// })
		}
	}

	if !foundAMD {
		fmt.Println("No active AMD GPUs detected via WMI.")
		// return nil, error.New("cannot found AMD GPU")
	}

	for _, eng := range m.engines {
		luid := helpers.ExtractLUID(eng.Name)

		if item, exists := gpuMap[luid]; exists {
			item.EngineLoad += eng.UtilizationPercentage
		}
	}

	for _, mem := range m.memories {
		luid := helpers.ExtractLUID(mem.Name)

		if item, exists := gpuMap[luid]; exists {
			// CurrentUsage is tracked by Windows in bytes -> convert to Megabytes
			item.Data.Mem_used += mem.CurrentUsage / (1024 * 1024)
		}
	}

	for _, item := range gpuMap {
		if item.EngineLoad > 100 {
			item.EngineLoad = 100
		}

		item.Data.Load = float64(item.EngineLoad)

		if item.Data.Mem_total > 0 {
			// Windows tracks global architecture limits inside memory performance tables.
			// Cap allocation counters at total physical capacity bounds to maintain integrity.
			if item.Data.Mem_used > item.Data.Mem_total {
				item.Data.Mem_used = item.Data.Mem_total
			}

			item.Data.Mem_percent = (float64(item.Data.Mem_used) / float64(item.Data.Mem_total)) * 100
		}

		gpus = append(gpus, item.Data)
	}

	return gpus, nil
}

// Fetch Static Hardware Info
func (m *Monitor) getAdapter() {
	query := "SELECT Name, PNPDeviceID, AdapterRAM FROM Win32_VideoController"
	err := wmi.Query(query, &m.controllers)
	
	if err != nil {
		fmt.Printf("Failed to query WMI Adapter: %v\n", err)
		return
	}
}

// Fetch Live Utilization Engines
func (m *Monitor) getEngines() {
	query := "SELECT Name, UtilizationPercentage FROM Win32_PerfFormattedData_GPUPerformanceCounters_GPUEngine"
	err := wmi.Query(query, &m.engines)
	
	if err != nil {
		fmt.Printf("Failed to query WMI Engines: %v\n", err)
		return
	}
}

// Fetch OS-allocated Virtual Memory tracks
func (m *Monitor) getMemories() {
	query := "SELECT Name, CurrentUsage FROM Win32_PerfFormattedData_GPUPerformanceCounters_GPUMemory"
	err := wmi.Query(query, &m.memories)
	
	if err != nil {
		fmt.Printf("Failed to query WMI Memories: %v\n", err)
		return
	}
}

