package functions

import (
	"fmt"
	"log"

	g "github.com/gosnmp/gosnmp"
)

func Walk(config *g.GoSNMP, root string) {
	err := config.BulkWalk(root, func(pdu g.SnmpPDU) error {
		// print OID and its value type
		fmt.Printf("%s = ", pdu.Name)

		switch pdu.Type {
            case g.OctetString:
                fmt.Printf("STRING: %s\n", string(pdu.Value.([]byte)))
            default:
                fmt.Printf("TYPE %d: %v\n", pdu.Type, pdu.Value)
		}
		return nil // Continue walking
	})

	if err != nil {
		log.Fatalf("BulkWalk() err: %v", err)
	}
}
