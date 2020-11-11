package wrappers

import (
	"reflect"

	"github.com/golang/protobuf/ptypes/wrappers"
)

func ParseDoubleValue(v interface{}) (value float64, valid bool) {
	if reflect.ValueOf(v).IsNil() == false {
		val := v.(*wrappers.DoubleValue)
		return val.Value, true
	}
	return 0, false

}

func ParseFloatValue(v interface{}) (value float32, valid bool) {
	if reflect.ValueOf(v).IsNil() == false {
		val := v.(*wrappers.FloatValue)
		return val.Value, true
	}
	return 0, false

}

func ParseInt64Value(v interface{}) (value int64, valid bool) {
	if reflect.ValueOf(v).IsNil() == false {
		val := v.(*wrappers.Int64Value)
		return val.Value, true
	}
	return 0, false

}

func ParseUInt64Value(v interface{}) (value uint64, valid bool) {
	if reflect.ValueOf(v).IsNil() == false {
		val := v.(*wrappers.UInt64Value)
		return val.Value, true
	}
	return 0, false

}

func ParseInt32Value(v interface{}) (value int32, valid bool) {
	if reflect.ValueOf(v).IsNil() == false {
		val := v.(*wrappers.Int32Value)
		return val.Value, true
	}
	return 0, false

}

func ParseUInt32Value(v interface{}) (value uint32, valid bool) {
	if reflect.ValueOf(v).IsNil() == false {
		val := v.(*wrappers.UInt32Value)
		return val.Value, true
	}
	return 0, false

}

func ParseBoolValue(v interface{}) (value bool, valid bool) {
	if reflect.ValueOf(v).IsNil() == false {
		val := v.(*wrappers.BoolValue)
		return val.Value, true
	}
	return false, false

}

func ParseStringValue(v interface{}) (value string, valid bool) {
	if reflect.ValueOf(v).IsNil() == false {
		val := v.(*wrappers.StringValue)
		return val.Value, true
	}
	return "", false

}

func ParseBytesValue(v interface{}) (value []byte, valid bool) {
	if reflect.ValueOf(v).IsNil() == false {
		val := v.(*wrappers.BytesValue)
		return val.Value, true
	}
	return nil, false

}
