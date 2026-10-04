package models

import "pc-monitoring/models/config"

type PageHeader struct {
	Title       string
	Description string
}

type PageData struct {
    Header     PageHeader
    Info       *Response
    Companies  []config.Company
}
