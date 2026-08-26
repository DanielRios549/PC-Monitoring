package helpers

import (
	"fmt"

	g "github.com/gosnmp/gosnmp"
)

func GetInfo(config *g.GoSNMP, oids []string) {
    var result *g.SnmpPacket
    var err error

    switch config.Version {
        case 0:
            result, err = config.Get(oids)
        default:
            result, err = config.GetBulk(oids, 1, 1)
    }

    if err != nil {
        fmt.Printf("Get() err: %v", err)
        return
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
