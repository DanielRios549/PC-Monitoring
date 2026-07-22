package helpers

import "math"

func RoundTo(val float64, places int) float64 {
	if places < 0 {
		return val
	}

	factor := math.Pow(10, float64(places))
	return math.Round(val * factor) / factor
}
