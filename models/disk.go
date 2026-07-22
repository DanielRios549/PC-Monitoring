package models

type DiskData struct {
	Device  string  `json:"device"`
	Mount   string  `json:"mount"`
	Total   float64 `json:"total"`
	Used    float64 `json:"used"`
	Percent float64 `json:"percent"`
}
