package models

type Response struct {
	CPU  []*CPUData  `json:"cpus"`
	RAM  []*RAMData  `json:"ram"`
	Disk []*DiskData `json:"disks"`
	GPU  []*GPUData  `json:"gpus"`
}
