package helpers

import (
	"fmt"
	"pc-monitoring/models"

	g "github.com/gosnmp/gosnmp"
)

var RootOID = "1.3.6.1"

func GetInfo(config *g.GoSNMP, options map[string][]string) []*models.OidRespose {
    var result *g.SnmpPacket
    var err error

    var items []*models.OidRespose

    for key, option := range options {
        getOid := option[1]

        if config.Version == 0 {
            getOid = option[0]
        }

        value := ""

        switch config.Version {
            case 0:
                result, err = config.Get([]string{RootOID + getOid})
            default:
                result, err = config.GetBulk([]string{RootOID + getOid}, 1, 1)
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

            items = append(items, &models.OidRespose{
                Name: key,
                Oid: getOid,
                Value: value,
            })
        }
    }

    return items
}
