package snmp

import (
	"log"

	g "github.com/gosnmp/gosnmp"
)

func WalkCount(config *g.GoSNMP, root string) int {
    count := 0

    err := config.BulkWalk(root, func(pdu g.SnmpPDU) error {
        count++
		return nil
    })

	if err != nil {
		log.Fatalf("Walk err: %v", err)
	}

    return count
}

type WalkTable map[string]string

func Walk(config *g.GoSNMP, root string) WalkTable {
    var table WalkTable

	err := config.BulkWalk(root, func(pdu g.SnmpPDU) error {
		// fmt.Printf("%s = ", pdu.Name)

		switch pdu.Type {
            case g.OctetString:
                // fmt.Printf("STRING: %s\n", string(pdu.Value.([]byte)))
                table[pdu.Name] = string(pdu.Value.([]byte))
            default:
                // fmt.Printf("TYPE %d: %v\n", pdu.Type, pdu.Value)
                table[pdu.Name] = g.ToBigInt(pdu.Value).String()
		}
		return nil // Continue walking
	})

	if err != nil {
		log.Fatalf("BulkWalk() err: %v", err)
	}

    return table
}
