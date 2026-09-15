package models

type Response struct {
	CPU  *CPUInfo  `json:"cpu"`
	RAM  *RAMData  `json:"ram"`
	Disk *DiskData `json:"disk"`
	GPU  *GPUData  `json:"gpu"`
}

type APResponse struct {
    Hostname         string  `json:"hostname"`
    Model            string  `json:"ap_model"`
    Version          string  `json:"version"`
    Devices          string  `json:"devices"`
}

type PrinterResponse struct {
    Hostname         string  `json:"hostname"`
    Model            string  `json:"printer_model"`
    Toner_Percent    float32 `json:"toner_percent"`
}
