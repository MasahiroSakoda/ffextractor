package util

import (
	"errors"
	"math"
	"testing"
	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/require"
)

func TestParseType(t *testing.T) {
	cases := []struct {
		value  string
		expect Type
	}{
		{"{}", TypeObject},
		{"[]", TypeArray},
		{"1", TypeInteger},
		{"1.5", TypeNumber},
		{"false", TypeBoolean},
		{"true", TypeBoolean},
		{"2015-09-03T13:27:52Z", TypeString},
		{"", TypeString},
		{"Go to https://golang.org for more information", TypeString},
	}
	for i, c := range cases {
		got := ParseType([]byte(c.value))
		if c.expect != got {
			t.Errorf("case %d response mismatch. expected: %s, got: %s", i, c.expect, got)
			continue
		}
	}
}

func TestParseString(t *testing.T) {
	cases := []struct {
		input  []byte
		expect string
		err    error
	}{
		{[]byte("foo"), "foo", nil},
	}
	for i, c := range cases {
		value, got := ParseString(c.input)
		if value != c.expect {
			t.Errorf("case %d value mismatch. expected: %s, got: %s", i, c.expect, value)
		}
		if c.err != got {
			t.Errorf("case %d error mismatch. expected: %s, got: %s", i, c.err, got)
		}
	}
}

func TestParseNumber(t *testing.T) {
	cases := []struct {
		input  []byte
		expect float64
		err    error
	}{
		{[]byte("1234567890"), float64(1234567890), nil},
		{[]byte("12345.67890"), float64(12345.67890), nil},
		{[]byte("-12345.67890"), float64(-12345.67890), nil},
		{[]byte("1.797693134862315708145274237317043567981e+308"), math.MaxFloat64, nil},
		{[]byte("2e+308"), math.Inf(0), errors.New(`strconv.ParseFloat: parsing "2e+308": value out of range`)},
		{[]byte("4.940656458412465441765687928682213723651e-324"), math.SmallestNonzeroFloat64, nil},
		{[]byte("1.940e-324"), float64(0), nil},
	}
	for i, c := range cases {
		value, got := ParseNumber(c.input)
		if value != c.expect {
			t.Errorf("case %d value mismatch. expected: %e, got: %e", i, c.expect, value)
		}
		if got != nil {
			if c.err != nil && got.Error() != c.err.Error() {
				t.Errorf("case %d error mismatch. expected: %s, got: %s", i, c.err, got)
			}
		}
	}
}

func TestParseInteger(t *testing.T) {
	cases := []struct {
		input  []byte
		expect int64
		err    error
	}{
		{[]byte(""), 0, errors.New(`strconv.ParseInt: parsing "": invalid syntax`)},
		{[]byte("9223372036854775807"), math.MaxInt64, nil},
		{[]byte("9223372036854775808"), math.MaxInt64, errors.New(`strconv.ParseInt: parsing "9223372036854775808": value out of range`)},
		{[]byte("-9223372036854775808"), math.MinInt64, nil},
		{[]byte("-9223372036854775809"), math.MinInt64, errors.New(`strconv.ParseInt: parsing "-9223372036854775809": value out of range`)},
		{[]byte("1234567890"), int64(1234567890), nil},
		{[]byte("12345.67890"), 0, errors.New(`strconv.ParseInt: parsing "12345.67890": invalid syntax`)},
		{[]byte("-12345.67890"), 0, errors.New(`strconv.ParseInt: parsing "-12345.67890": invalid syntax`)},
	}
	for i, c := range cases {
		value, got := ParseInteger(c.input)
		if value != c.expect {
			t.Errorf("case %d value mismatch. expected: %d, got: %d", i, c.expect, value)
		}
		if got != nil {
			if c.err != nil && got.Error() != c.err.Error() {
				t.Errorf("case %d error mismatch. expected: %s, got: %s", i, c.err, got)
			}
		}
	}
}

func TestParseBoolean(t *testing.T) {
	cases := []struct {
		input  []byte
		expect bool
		err    error
	}{}
	for i, c := range cases {
		value, got := ParseBoolean(c.input)
		if value != c.expect {
			t.Errorf("case %d value mismatch. expected: %t, got: %t", i, c.expect, value)
		}
		if c.err != got {
			t.Errorf("case %d error mismatch. expected: %s, got: %s", i, c.err, got)
		}
	}
}

func TestIsInteger(t *testing.T) {
	cases := []struct {
		b      []byte
		expect bool
	}{
		{[]byte(""), false},
		{[]byte("1"), true},
		{[]byte("367890"), true},
		{[]byte("1.2"), false},
		{[]byte("foo"), false},
		{[]byte("9223372036854775808"), true},
		{[]byte("890oasdfg dfgh89"), false},
		{[]byte("[123]"), false},
	}
	for _, c := range cases {
		got := IsInteger(c.b)
		if got != c.expect {
			t.Errorf("case IsInteger: %s - expected: '%t', got: '%t'", c.b, c.expect, got)
		}
	}
}

func TestIsBoolean(t *testing.T) {
	tests := []struct {
		name     string
		value    []byte
		expected bool
	}{
		{ name: "with positive integer value", value: []byte("10"),    expected: false },
		{ name: "with negative integer value", value: []byte("-10"),   expected: false },
		{ name: "with positive float value",   value: []byte("3.14"),  expected: false },
		{ name: "with negative float value",   value: []byte("-3.14"), expected: false },
		{ name: "with bool value",             value: []byte("false"), expected: true },
		{ name: "with text",                   value: []byte("blah"),  expected: false },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsBoolean(tt.value)
			assert.Equal(t, tt.expected, got)
		})
	}
}
