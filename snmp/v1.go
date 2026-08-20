package snmp

import (
	"fmt"
	"log"
	"os"

	g "github.com/gosnmp/gosnmp"
)

func V1() {
	// Default is a pointer to a GoSNMP struct that contains sensible defaults
	// eg port 161, community public, etc
	envTarget := os.Getenv("GOSNMP_TARGET")

	if len(envTarget) == 0 {
		log.Fatalf("environment variable not set: GOSNMP_TARGET")
	}

	g.Default.Target = envTarget
	err := g.Default.Connect()

	if err != nil {
		log.Fatalf("Connect() err: %v", err)
	}

	defer g.Default.Conn.Close()

	oids := []string{"1.3.6.1.2.1.1.4.0", "1.3.6.1.2.1.1.7.0"}
	result, err2 := g.Default.Get(oids) // Get() accepts up to g.MAX_OIDS
	if err2 != nil {
		log.Fatalf("Get() err: %v", err2)
	}

	for i, variable := range result.Variables {
		fmt.Printf("%d: oid: %s ", i, variable.Name)

		// the Value of each variable returned by Get() implements
		// interface{}. You could do a type switch...
		switch variable.Type {
		case g.OctetString:
			bytes := variable.Value.([]byte)
			fmt.Printf("string: %s\n", string(bytes))
		default:
			// ... or often you're just interested in numeric values.
			// ToBigInt() will return the Value as a BigInt, for plugging
			// into your calculations.
			fmt.Printf("number: %d\n", g.ToBigInt(variable.Value))
		}
	}
}
