package models

type SpeedTypes struct {
	Current string `json:"current"`
	Total   string `json:"total"`
}

type Speed struct {
	Read  SpeedTypes `json:"read"`
	Write SpeedTypes `json:"write"`
}

type DiskData struct {
	Device  string  `json:"device"`
	Mount   string  `json:"mount"`
	Total   float64 `json:"total"`
	Used    float64 `json:"used"`
	Percent float64 `json:"percent"`
}
