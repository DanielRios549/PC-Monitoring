package snmp

import (
	"errors"
	"fmt"
	"pc-monitoring/helpers"
	"pc-monitoring/models"
	"time"

	g "github.com/gosnmp/gosnmp"
)

func V2(ip string) ([]*models.OidRespose, error) {
	params := &g.GoSNMP{
		Target:    ip,
		Port:      161,
		Community: "public",
		Version:   g.Version2c,
		Timeout:   time.Duration(2) * time.Second,
		// Logger:    g.NewLogger(log.New(os.Stdout, "", 0)),
	}

	err := params.Connect()

	if err != nil {
		// fmt.Printf("Connect() error: %v", err)
        return nil, errors.New("printer is offline")
	}

	defer func() {
		err := params.Conn.Close()

		if err != nil {
			fmt.Printf("Cannot Close Connection: %v", err)
		}
	}()

	info := helpers.GetInfo(params, Options)

    return info, nil
}
