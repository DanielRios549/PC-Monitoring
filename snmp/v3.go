package snmp

import (
	"errors"
	"fmt"
	"time"
	"pc-monitoring/helpers"

	g "github.com/gosnmp/gosnmp"
)

func V3(ip, context, user, pass, privpass string) error {
	params := &g.GoSNMP{
		Target:        ip,
		Port:          161,
		Version:       g.Version3,
		SecurityModel: g.UserSecurityModel,
		MsgFlags:      g.AuthPriv,
		Timeout:       time.Duration(30) * time.Second,
		ContextName:  context,
		SecurityParameters: &g.UsmSecurityParameters{
			UserName: user,
			AuthenticationProtocol:   g.MD5,
			AuthenticationPassphrase: pass,
			PrivacyProtocol:          g.DES,
			PrivacyPassphrase:        privpass,
		},
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

	// Function handles for collecting metrics on query latencies.
	// var sent time.Time

	// params.OnSent = func(_ *g.GoSNMP) {
	// 	sent = time.Now()
	// }
	// params.OnRecv = func(_ *g.GoSNMP) {
	// 	log.Println("Query latency in seconds:", time.Since(sent).Seconds())
	// }
	
	// OIDs
	// MIB Root: 1.3.6.1.2.1.43
	// Toner:    1.3.6.1.2.1.43.11.1.1
	// Paper:    1.3.6.1.2.1.43.8.2.1

	rootOID    := "1.3.6.1"
	// rootPages  := ".2.1.43.8"
	// rootToner  := ".2.1.43.11"

	oids := []string{
		rootOID + ".2.1.1.4.0",        // Host Name
        rootOID + ".2.1.25.3.2.1.3",   // Printer Model
        rootOID + ".2.1.43.11.1.1.9",  // Toner Current Level
        rootOID + ".2.1.43.11.1.1.8",  // Toner Maximum Level
	}

	helpers.GetInfo(params, oids)
	// functions.Walk(params, rootOID + rootPages)

    return nil
}
