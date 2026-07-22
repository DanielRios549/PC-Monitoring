package models

type GPUInfo struct {
	Name        string
	Vendor      string // "Nvidia" "AMD" or "Intel"
	Description string
}

type GPUData struct {
	Name        string  `json:"name"`
	Load        float64 `json:"load"`      // This is the 'usage' in your dashboard loop
	Mem_total   int     `json:"mem_total"` // 1024 based mb
	Mem_used    int     `json:"mem_used"`
	Mem_percent float64 `json:"mem_percent"`
}
