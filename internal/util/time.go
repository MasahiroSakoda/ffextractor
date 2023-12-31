// Package util provides function for utility
package util

import (
	"fmt"
	"math"
	"strings"
	"strconv"
)

// ConvertToSeconds converts a string like 00:00:11.35 into seconds (11.35)
func ConvertToSeconds(s string) (seconds float64, err error) {
	s = strings.TrimSpace(s)
	parts := strings.Split(s, ":")
	multipliers := []float64{60 * 60, 60, 1}
	if len(parts) == 2 {
		multipliers = []float64{60, 1, 1}
	} else if len(parts) == 1 {
		multipliers = []float64{1, 1, 1}
	}
	for i, part := range parts {
		var partf float64
		partf, err = strconv.ParseFloat(part, 64)
		if err != nil {
			return
		}
		seconds += partf * multipliers[i]
	}
	return
}

// SecondsToString seconds like 80 to a string like 00:01:20.00
func SecondsToString(seconds float64) string {
	hours := math.Floor(seconds / 3600)
	seconds = seconds - hours*3600

	minutes := math.Floor(seconds / 60)
	seconds = seconds - minutes*60

	s := fmt.Sprintf("%02d:%02d:%02.4f", int(hours), int(minutes), seconds)
	if seconds < 10 {
		s = fmt.Sprintf("%02d:%02d:0%2.4f", int(hours), int(minutes), seconds)
	}
	for i := 0; i < 3; i++ {
		s = strings.TrimSuffix(s, "0")
	}
	return s
}
