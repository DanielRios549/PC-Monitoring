package snmp

import (
	"fmt"
	"pc-monitoring/models/config"
	"strconv"
	"strings"

	g "github.com/gosnmp/gosnmp"
)

var RootOID    = "1.3.6.1"
var RootVendor = "4.1"

func GetInfo(snmp *g.GoSNMP, options map[string][]string) []*config.Oid {
    var result *g.SnmpPacket
    var err error

    var items []*config.Oid

    // TODO: Add parameter
    // TODO: Get simple information only to show items in a list
    simple := false

    if simple {
        fmt.Println("WORK IN POGRESS")
    } else {
        for key, option := range options {
            getOid := option[1]

            if snmp.Version == 0 {
                getOid = option[0]
            }

            value := "N/A"

            if strings.HasSuffix(key, "_count") {
                count := WalkCount(snmp, RootOID + getOid)
                value = strconv.Itoa(count)
            } else if strings.HasSuffix(key, "_get") {
                result, err := snmp.Get([]string{RootOID + getOid})

                if err != nil {
                    fmt.Printf("Get() Vendor err: %v\n", err)
                }

                variable := result.Variables[0]

                fullValue := variable.Value.(string)

                splitted := strings.SplitAfter(
                    fullValue,
                    fmt.Sprintf(".%s.%s.", RootOID, RootVendor),
                )

                if len(splitted) > 1 {
                    value = strings.Split(splitted[1], ".")[0]
                } else {
                    value = splitted[0]
                }
            } else if strings.HasSuffix(key, "_next") {
                result, err := snmp.GetNext([]string{RootOID + getOid})

                if err != nil {
                    fmt.Printf("GetNext() Vendor err: %v\n", err)
                }

                variable := result.Variables[0]
                value = variable.Value.(string)
            } else {
                switch snmp.Version {
                    case 0:
                        result, err = snmp.Get([]string{RootOID + getOid})
                    default:
                        result, err = snmp.GetBulk([]string{RootOID + getOid}, 1, 1)
                }

                if err != nil {
                    fmt.Printf("Get() err: %v\n", err)
                } else {
                    if len(result.Variables) < 1 {
                        fmt.Printf("Variables Empty: %v\n", err)
                    } else {
                        variable := result.Variables[0]

                        // the Value of each variable returned by Get() implements
                        // interface{}. You could do a type switch...

                        switch variable.Type {
                            case g.OctetString:
                                value = string(variable.Value.([]byte))
                                // fmt.Printf("%s: %s\n", key, value)
                            default:
                                // ... or often you're just interested in numeric values.
                                // ToBigInt() will return the Value as a BigInt, for plugging
                                // into your calculations.
                                currentValue := g.ToBigInt(variable.Value)
                                value = currentValue.String()
                                // fmt.Printf("%s: %d\n", key, currentValue)
                        }
                    }
                }
            }

            items = append(items, &config.Oid{
                Name: key,
                Oid: getOid,
                Value: value,
            })
        }
    }

    return items
}
