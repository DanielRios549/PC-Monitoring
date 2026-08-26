package snmp

import (
	"errors"
	"fmt"
	"time"
	"pc-monitoring/helpers"

	g "github.com/gosnmp/gosnmp"
)

func V2(ip string) error {
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

	oids := []string{"1.3.6.1.2.1.1.5.0"}

	helpers.GetInfo(params, oids)

    return nil
}
