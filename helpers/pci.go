package helpers

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func isGPU(path string) bool {
	class := ReadString(filepath.Join(path, "class"))

	// 0x030000 = VGA controller
	// 0x030200 = 3D controller
	return strings.HasPrefix(class, "0x03")
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

	return result
}

