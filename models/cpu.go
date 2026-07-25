package models

type CPUInfo struct {
	Name      string `json:"name"`
	CoreCount int    `json:"threads"`
}

type CPUData struct {
	Core    int     `json:"core"`
	Usage   float64 `json:"usage"`
	Color   string  `json:"color"`
}
