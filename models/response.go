package models

type Response struct {
	CPU  *CPUInfo  `json:"cpu"`
	RAM  *RAMData  `json:"ram"`
	Disk *DiskData `json:"disk"`
	GPU  *GPUData  `json:"gpu"`
}


