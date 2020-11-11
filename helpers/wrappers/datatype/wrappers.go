package datatype

// BoolValue get pointer value
func BoolValue(v *bool) bool {
	if v != nil {
		return *v
	}
	return false
}

// StringValue get pointer value
func StringValue(v *string) string {
	if v != nil {
		return *v
	}
	return ""
}

// IntValue get pointer value
func IntValue(v *int) int {
	if v != nil {
		return *v
	}
	return 0
}

// Int8Value get pointer value
func Int8Value(v *int8) int8 {
	if v != nil {
		return *v
	}
	return 0
}

// Int16Value get pointer value
func Int16Value(v *int16) int16 {
	if v != nil {
		return *v
	}
	return 0
}

// Int32Value get pointer value
func Int32Value(v *int32) int32 {
	if v != nil {
		return *v
	}
	return 0
}

// Int64Value get pointer value
func Int64Value(v *int64) int64 {
	if v != nil {
		return *v
	}
	return 0
}

// UIntValue get pointer value
func UIntValue(v *uint) uint {
	if v != nil {
		return *v
	}
	return 0
}

// UInt8Value get pointer value
func UInt8Value(v *uint8) uint8 {
	if v != nil {
		return *v
	}
	return 0
}

// UInt16Value get pointer value
func UInt16Value(v *uint16) uint16 {
	if v != nil {
		return *v
	}
	return 0
}

// UInt32Value get pointer value
func UInt32Value(v *uint32) uint32 {
	if v != nil {
		return *v
	}
	return 0
}

// UInt64Value get pointer value
func UInt64Value(v *uint64) uint64 {
	if v != nil {
		return *v
	}
	return 0
}

// Float32Value get pointer value
func Float32Value(v *float32) float32 {
	if v != nil {
		return *v
	}
	return 0
}

// Float64Value get pointer value
func Float64Value(v *float64) float64 {
	if v != nil {
		return *v
	}
	return 0
}

// UIntPtrValue get pointer value
func UIntPtrValue(v *uintptr) uintptr {
	if v != nil {
		return *v
	}
	return 0
}

// Complex64Value get pointer value
func Complex64Value(v *complex64) complex64 {
	if v != nil {
		return *v
	}
	return 0
}

// Complex128Value get pointer value
func Complex128Value(v *complex128) complex128 {
	if v != nil {
		return *v
	}
	return 0
}
