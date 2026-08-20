package snmp

import (
	"log"
	"os"
	"pc-monitoring/functions"

	// "strconv"
	"time"

	g "github.com/gosnmp/gosnmp"
)

func V3() {
	// get Target and Port from environment
	envTarget := os.Getenv("GOSNMP_TARGET")
	envPort := os.Getenv("GOSNMP_PORT")

	if len(envTarget) == 0 {
		log.Fatalf("environment variable not set: GOSNMP_TARGET")
	}
	if len(envPort) == 0 {
		log.Fatalf("environment variable not set: GOSNMP_PORT")
	}

	// port, _ := strconv.ParseUint(envPort, 10, 16)

	// build our own GoSNMP struct, rather than using g.Default
	params := &g.GoSNMP{
		Target:        "10.10.6.200", // 200, 142, 167 envTarget,
		Port:          161, // uint16(port),
		Version:       g.Version3,
		SecurityModel: g.UserSecurityModel,
		MsgFlags:      g.AuthPriv,
		Timeout:       time.Duration(30) * time.Second,
		// ContextName: "Jetdirect",
		SecurityParameters: &g.UsmSecurityParameters{
			UserName: os.Getenv("GOSNMP_USER"),
			AuthenticationProtocol:   g.MD5,
			AuthenticationPassphrase: os.Getenv("GOSNMP_PASS"),
			PrivacyProtocol:          g.DES,
			PrivacyPassphrase:        os.Getenv("GOSNMP_PRIVPASS"),
		},
	}

	err := params.Connect()

	if err != nil {
		log.Fatalf("Connect() err: %v", err)
	}

	defer func(){
		err := params.Conn.Close()

		if err != nil {
			log.Fatalf("Cannot Close Connection: %v", err)
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
	rootPages  := ".2.1.43.8"
	// rootToner  := ".2.1.43.11"

	// oids := []string{
	// 	rootOID + ".2.1.1.4.0",
	// 	rootOID + ".2.1.43.11.1.1.6",
	// 	rootOID + ".2.1.25.3.5.1.1.1",
	// 	rootOID + ".2.1.25.3.5.1.2.1",
	// 	rootOID + ".2.1.43.10.2.1.4",
	// 	rootOID + ".2.1.43.11.1.1.9",
	// }

	// functions.GetInfo(params, oids)
	functions.Walk(params, rootOID + rootPages)
}
