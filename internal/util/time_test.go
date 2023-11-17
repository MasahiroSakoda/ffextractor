package util

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestConvertToSeconds(t *testing.T) {
	tests := []struct{
		name     string
		time     string
		expected float64
	}{
		{ name: "", time: "00:11.5",    expected: 11.5 },
		{ name: "", time: "00:00:11.5", expected: 11.5 },
		{ name: "", time: "00:01:11.5", expected: 71.5 },
		{ name: "", time: "01:01:11.5", expected: 3671.5 },
		{ name: "", time: "11.5",       expected: 11.5 },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seconds, err := ConvertToSeconds(tt.time)
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, seconds)
		})
	}
}

func TestSecondsToString(t *testing.T) {
	tests  := []struct{
		name     string
		seconds  float64
		expected string
	}{
		{ name: "", seconds: 64.23,   expected: "00:01:04.23" },
		{ name: "", seconds: 90.23,   expected: "00:01:30.23" },
		{ name: "", seconds: 180.23,  expected: "00:03:00.23" },
		{ name: "", seconds: 3690.23, expected: "01:01:30.23" },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := SecondsToString(tt.seconds)
			assert.Equal(t, tt.expected, str)
		})
	}
}
