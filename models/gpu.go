package models

type GPUMonitor interface{
	Stats() ([]*GPUData, error)
	CountDevices() int8
	Close()
	// getAdapter()
	// getEngines()
	// getMemories()
}

type GPUInfo struct {
	Name        string
	Vendor      string // "Nvidia" "AMD" or "Intel"
	Description string
}

type GPUData struct {
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Load        float64 `json:"load"`      // GPU usage
	Mem_total   uint64  `json:"mem_total"` // 1024 based mb
	Mem_used    uint64  `json:"mem_used"`
	Mem_percent float64 `json:"mem_percent"`
}
