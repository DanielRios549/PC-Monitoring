package models



type RAMData struct {
	Total_gb  float64 `json:"total_gb"`
	Used_gb   float64 `json:"used_gb"`
	Percent   float64 `json:"percent"`
}
