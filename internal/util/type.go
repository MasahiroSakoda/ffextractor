// Package util provides function for utility
package util

import (
	"bytes"
	"strconv"
)

// Type is a type of data, these types follow JSON type primitives,
// with an added type for whole numbers (integer)
type Type uint8

const (
	// TypeUnknown is the default datatype, making for easier
	// errors when a datatype is expected
	TypeUnknown Type = iota
	// TypeNull specifies the null type
	TypeNull
	// TypeInteger specifies whole numbers
	TypeInteger
	// TypeNumber specifies numbers with decimal value
	TypeNumber
	// TypeBoolean species true/false values
	TypeBoolean
	// TypeString specifies text values
	TypeString
	// TypeObject maps string keys to values
	TypeObject
	// TypeArray is an ordered list of values
	TypeArray
	// TypeBytes is an ordered slice of bytes
	TypeBytes
)

// ParseType examines a slice of bytes & attempts to determine
// it's type, starting with the more specific possible types, then falling
// back to more general types. ParseType always returns a type
func ParseType(value []byte) Type {
	if IsBoolean(value) {
		return TypeBoolean
	}
	if bytes.Equal(value, []byte("null")) {
		return TypeNull
	}
	for _, b := range value {
		switch b {
		case '"':
			return TypeString
		case '-', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0', 'e':
			if _, e := strconv.ParseFloat(string(value), 32); e != nil {
				return TypeString
			}
			if IsInteger(value) {
				return TypeInteger
			}
			return TypeNumber
		case '{':
			return TypeObject
		case '[':
			return TypeArray
		case ' ', '\n':
			continue
		default:
			return TypeString
		}
	}

	// assume a string? sure.
	return TypeString
}

// String satsfies the stringer interface
func (dt Type) String() string {
	s, ok := map[Type]string{
		TypeUnknown: "",
		TypeString:  "string",
		TypeInteger: "integer",
		TypeNumber:  "number",
		TypeBoolean: "boolean",
		TypeObject:  "object",
		TypeArray:   "array",
		TypeNull:    "null",
	}[dt]

	if !ok {
		return ""
	}

	return s
}

// ParseString converts raw bytes to a string value
func ParseString(value []byte) (string, error) {
	return string(value), nil
}

// ParseNumber converts raw bytes to a float64 value
func ParseNumber(value []byte) (float64, error) {
	return strconv.ParseFloat(string(value), 64)
}

// ParseInteger converts raw bytes to a int64 value
func ParseInteger(value []byte) (int64, error) {
	return strconv.ParseInt(string(value), 10, 64)
}

// ParseBoolean converts raw bytes to a bool value
func ParseBoolean(value []byte) (bool, error) {
	return strconv.ParseBool(string(value))
}

// IsInteger checks if a slice of bytes is an integer
func IsInteger(value []byte) bool {
	if len(value) == 0 {
		return false
	}
	if value[0] == '[' || value[0] == '{' || !bytes.ContainsAny(value[0:1], "-+0123456789") {
		return false
	}
	if _, err := ParseInteger(value); err == nil || err.(*strconv.NumError).Err == strconv.ErrRange {
		return true
	}
	return false
}

// IsFloat checks if a slice of bytes is a float value
func IsFloat(value []byte) bool {
	if len(value) == 0 {
		return false
	}
	if value[0] == '[' || value[0] == '{' || !bytes.ContainsAny(value[0:1], "-+0123456789") {
		return false
	}
	if _, err := ParseNumber(value); err == nil || err.(*strconv.NumError).Err == strconv.ErrRange {
		return true
	}
	return false
}

// IsBoolean checks if a slice of bytes is a boolean value
func IsBoolean(value []byte) bool {
	switch string(value) {
	case "t", "f", "T", "F", "true", "false", "TRUE", "FALSE", "True", "False":
		return true
	default:
		return false
	}
}
