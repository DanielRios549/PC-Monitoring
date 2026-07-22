package monitors

import "pc-monitoring/models"

// TODO: Add Real GPU data

// GPU details
func GPU() []*models.GPUData {
	gpu_list := []*models.GPUData{}

	for i := 0; i < 1; i++ { // Placeholder for loop if multiple GPUs found
		gpu_list = append(gpu_list, &models.GPUData{
			Name:        "GPU",
			Load:        0.0, // TODO: Replace with actual logic or NVML call
			Mem_total:   8192, 
			Mem_used:    0,
			Mem_percent: 0.0,
		})
	}

	return gpu_list
}
