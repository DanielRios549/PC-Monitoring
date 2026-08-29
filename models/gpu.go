package models

import "time"

type Vendor string

const (
    NVIDIA Vendor = "Nvidia"
    AMD    Vendor = "AMD"
    Intel  Vendor = "Intel"
    Apple  Vendor = "Apple"
)

type GPUMonitor interface {
    Refresh() ([]*GPUData, error)
	CountDevices() int
	Close() error
	// getAdapter()
	// getEngines()
	// getMemories()
}

type GPUInfo struct {
	Name        string
	Vendor      Vendor
	Description string
}

type GPUData struct {
	ID          string     `json:"id"`
	Vendor      Vendor     `json:"vendor"`
	Name        string     `json:"name"`
	Driver      string     `json:"driver"`
	Type        string     `json:"type"`
	Load        float64    `json:"load"`
	Mem_total   uint64     `json:"mem_total"`
	Mem_used    uint64     `json:"mem_used"`
	Mem_percent float64    `json:"mem_percent"`
	Mem_free    uint64     `json:"mem_free"`
    Temperature float64    `json:"temperature"`
    PowerUsage  float64    `json:"power_usage"`
	PowerLimit  float64    `json:"power_limit"`
	CoreClock   uint64     `json:"core_clock"`
	MemoryClock uint64     `json:"memory_clock"`
	FanSpeed    float64    `json:"fan_speed"`
	BusID       string     `json:"bus_id"`
    Timestamp   time.Time  `json:"timestamp"`
}

