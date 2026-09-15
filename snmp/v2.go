package snmp

import (
	"errors"
	"fmt"
	// "pc-monitoring/models"
	"pc-monitoring/models/config"
	"pc-monitoring/snmp/oid"
	"time"

	g "github.com/gosnmp/gosnmp"
)

func V2(ip string, isAP bool) ([]*config.Oid, error) {
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

    oids := oid.PrinterOptions

    if isAP {
        oids = oid.APOptions
    }

	info := GetInfo(params, oids)

    return info, nil
}
