package models

type Response struct {
	CPU  *CPUInfo  `json:"cpu"`
	RAM  *RAMData  `json:"ram"`
	Disk *DiskData `json:"disk"`
	GPU  *GPUData  `json:"gpu"`
}

type OidRespose struct {
    Name  string  `json:"name"`
    Iod   string  `json:"oid"`
    Value string  `json:"value"`
}
