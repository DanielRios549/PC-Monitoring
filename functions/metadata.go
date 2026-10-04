package functions

import (
	"pc-monitoring/models"
	"pc-monitoring/monitors"
)

func PageData(title, description string) models.PageData {
    cpu := monitors.CPUInfo()

    header := models.PageHeader{
    	Title:       title,
    	Description: description,
    }

    info := &models.Response{
    	CPU: cpu,
    }

    return models.PageData{
    	Header:     header,
    	Info:       info,
        Companies:  LoadCompanies(),
    }
}
