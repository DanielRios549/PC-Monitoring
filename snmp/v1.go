package snmp

import (
	"errors"
	"fmt"
	"time"
	"pc-monitoring/helpers"

	g "github.com/gosnmp/gosnmp"
)

func V1(ip string) error {
	params := &g.GoSNMP{
		Target:    ip,
		Port:      161,
		Community: "public",
		Version:   g.Version1,
		Timeout:   time.Duration(2) * time.Second,
		// Logger:    g.NewLogger(log.New(os.Stdout, "", 0)),
	}

	err := params.Connect()

	if err != nil {
		// fmt.Printf("Connect() error: %v", err)
        return errors.New("printer is offline")
	}

	defer func() {
		err := params.Conn.Close()

		if err != nil {
			fmt.Printf("Cannot Close Connection: %v", err)
		}
	}()

    rootOID    := "1.3.6.1"
    
    // V1 need to use the exact OIDs since it has no GetBulk() support
	oids := []string{
		rootOID + ".2.1.1.5.0",          // Host Name
        rootOID + ".2.1.25.3.2.1.3.1",   // Printer Model
	}
	
    helpers.GetInfo(params, oids)

    return nil
}
