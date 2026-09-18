package functions

import (
	"fmt"
	"io"
	"net/http"
)

// TODO: Save Items on Disk
func GetVendorMac(mac string) string {
	mvAPI := "https://api.macvendors.com/" + mac
	resp, _ := http.Get(mvAPI)
	body, err := io.ReadAll(resp.Body)

	if err != nil {
        defer func() {
            err := resp.Body.Close()
            
            if err != nil {
                fmt.Printf("Get() Vendor err: %v\n", err)
            }
        }()
	}

	return string(body)
}
