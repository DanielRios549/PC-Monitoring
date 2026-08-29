package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	g "github.com/gosnmp/gosnmp"
)

var RootOID = "1.3.6.1"

func isGPU(path string) bool {
	class := ReadString(filepath.Join(path, "class"))

	// 0x030000 = VGA controller
	// 0x030200 = 3D controller
	return strings.HasPrefix(class, "0x03")
}

func GetInfo(config *g.GoSNMP, options map[string][]string) {
    var result *g.SnmpPacket
    var err error
    var oids []string

    for _, option := range options {
        getOid := option[1]

        if config.Version == 0 {
            getOid = option[0]
        }

        oids = append(oids, RootOID + getOid)
    }

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

func ReadString(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(data))
}

func ReadUint(path string) uint64 {
	value := ReadString(path)

	n, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0
	}

	return n
}

func ReadFloat(path string) float64 {
	value := ReadString(path)

	n, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}

	return n
}

func DriverForPCI(path string) string {
	target, err := os.Readlink(
		filepath.Join(path, "driver"),
	)

	if err != nil {
		return ""
	}

	return filepath.Base(target)
}

func PciAddress(path string) string {
	return filepath.Base(path)
}

func FindPCIByVendor(vendor string) []string {
	paths, _ := filepath.Glob(
		"/sys/bus/pci/devices/*",
	)

	var result []string

	for _, path := range paths {
		if !isGPU(path) {
			continue
		}

		if ReadString(
			filepath.Join(path, "vendor"),
		) != vendor {
			continue
		}

		result = append(result, path)
	}

    print(len(result))

	return result
}
