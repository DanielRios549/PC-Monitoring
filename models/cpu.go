package models

type CPUData struct {
	Core    int     `json:"core"`
	Usage   float64 `json:"usage"`
	Color   string  `json:"color"`
}
