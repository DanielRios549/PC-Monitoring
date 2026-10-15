package functions

import (
	"pc-monitoring/models"
)

func PageData(title, description string) models.PageData {
    header := models.PageHeader{
    	Title:       title,
    	Description: description,
    }

    return models.PageData{
    	Header:     header,
        Companies:  LoadCompanies(),
    }
}
