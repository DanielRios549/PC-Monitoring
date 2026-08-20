package functions

import (
	"fmt"
	"log"

	g "github.com/gosnmp/gosnmp"
)

func GetInfo(config *g.GoSNMP, oids []string) {
	result, err := config.GetBulk(oids, 1, 1) // Get() accepts up to g.MAX_OIDS

	if err != nil {
		log.Fatalf("Get() err: %v", err)
	}

	for i, variable := range result.Variables {
		fmt.Printf("%d: oid: %s ", i, variable.Name)

		// the Value of each variable returned by Get() implements
		// interface{}. You could do a type switch...
		switch variable.Type {
			case g.OctetString:
				fmt.Printf("string: %s\n", string(variable.Value.([]byte)))
			default:
				// ... or often you're just interested in numeric values.
				// ToBigInt() will return the Value as a BigInt, for plugging
				// into your calculations.
				fmt.Printf("number: %d\n", g.ToBigInt(variable.Value))
		}
	}
}
