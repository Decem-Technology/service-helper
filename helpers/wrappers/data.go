package wrappers

import (
	"regexp"
	"strings"

	_struct "github.com/golang/protobuf/ptypes/struct"
	"github.com/golang/protobuf/ptypes/wrappers"
)

// IsNotEmptyString check is not empty string
func IsNotEmptyString(v string) bool {
	if v == "" {
		return false
	}
	return true
}

// IsNotEmptyFloat64 check is not empty string
func IsNotEmptyFloat64(v float64) bool {
	if v == 0 {
		return false
	}
	return true
}

// IsParseValid check parse data valid
func IsParseValid(v error) bool {
	if v != nil {
		return false
	}
	return true
}

// WrapInt32 wrap int
func WrapInt32(v *wrappers.Int32Value) int32 {
	if v != nil {
		return v.Value
	}
	return 0
}

// WrapInt64 wrap Int64
func WrapInt64(v *wrappers.Int64Value) int64 {
	if v != nil {
		return v.Value
	}
	return 0
}

// WrapString wrap String
func WrapString(v *wrappers.StringValue) string {
	if v != nil {
		return v.Value
	}
	return ""
}

// WrapFloat32 wrap Float32
func WrapFloat32(v *wrappers.DoubleValue) float32 {
	if v != nil {
		return float32(v.Value)
	}
	return 0
}

// WrapFloat64 wrap Float64
func WrapFloat64(v *wrappers.DoubleValue) float64 {
	if v != nil {
		return v.Value
	}
	return 0
}

// StringValue wrap ngValue
func StringValue(v string) *_struct.Value {
	return &_struct.Value{Kind: &_struct.Value_StringValue{StringValue: v}}
}

// NumberValue wrap erValue
func NumberValue(v float64) *_struct.Value {
	return &_struct.Value{Kind: &_struct.Value_NumberValue{NumberValue: v}}
}

var camel = regexp.MustCompile("(^[^A-Z]*|[A-Z]*)([A-Z][^A-Z]+|$)")

// SnakeCase convert string to snake case
func SnakeCase(s string) string {
	var a []string
	for _, sub := range camel.FindAllStringSubmatch(s, -1) {
		if sub[1] != "" {
			a = append(a, sub[1])
		}
		if sub[2] != "" {
			a = append(a, sub[2])
		}
	}
	return strings.ToLower(strings.Join(a, "_"))
}
