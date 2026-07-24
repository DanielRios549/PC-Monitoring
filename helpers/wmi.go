package helpers

import (
	"regexp"
	"strings"
)

// Matches patterns like luid_0x00000000_0x0000F612
// or LUID inside hardware paths
func ExtractLUID(input string) string {
	re := regexp.MustCompile(`(?i)luid_(0x[0-9a-f]+_0x[0-9a-f]+)`)
	matches := re.FindStringSubmatch(input)

	if len(matches) > 1 {
		return strings.ToLower(matches[1])
	}

	return ""
}
